package asset

import (
	"atlas-inventory/cash"
	"atlas-inventory/equipable"
	"atlas-inventory/kafka/message"
	"atlas-inventory/kafka/message/asset"
	asset2 "atlas-inventory/kafka/producer/asset"
	model2 "atlas-inventory/model"
	"atlas-inventory/pet"
	"atlas-inventory/stackable"
	"context"
	"errors"
	"github.com/Chronicle20/atlas-constants/inventory"
	"github.com/Chronicle20/atlas-model/model"
	tenant "github.com/Chronicle20/atlas-tenant"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"time"
)

type Processor struct {
	l                  logrus.FieldLogger
	ctx                context.Context
	db                 *gorm.DB
	cashProcessor      *cash.Processor
	stackableProcessor *stackable.Processor
	GetByCompartmentId func(uuid.UUID) ([]Model[any], error)
}

func NewProcessor(l logrus.FieldLogger, ctx context.Context, db *gorm.DB) *Processor {
	p := &Processor{
		l:                  l,
		ctx:                ctx,
		db:                 db,
		cashProcessor:      cash.NewProcessor(l, ctx),
		stackableProcessor: stackable.NewProcessor(l, ctx, db),
	}
	p.GetByCompartmentId = model2.CollapseProvider(p.ByCompartmentIdProvider)
	return p
}

func (p *Processor) WithTransaction(db *gorm.DB) *Processor {
	return &Processor{
		l:                  p.l,
		ctx:                p.ctx,
		db:                 db,
		cashProcessor:      p.cashProcessor,
		stackableProcessor: p.stackableProcessor,
		GetByCompartmentId: p.GetByCompartmentId,
	}
}

func (p *Processor) ByCompartmentIdProvider(compartmentId uuid.UUID) model.Provider[[]Model[any]] {
	t := tenant.MustFromContext(p.ctx)
	ap := model.SliceMap(Make)(getByCompartmentId(t.Id(), compartmentId)(p.db))(model.ParallelMap())
	return model.SliceMap(p.DecorateAsset(compartmentId))(ap)(model.ParallelMap())
}

func (p *Processor) DecorateAsset(compartmentId uuid.UUID) func(m Model[any]) (Model[any], error) {
	return func(m Model[any]) (Model[any], error) {
		var decorator model.Transformer[Model[any], Model[any]]
		if m.IsEquipable() {
			decorator = p.DecorateEquipable
		} else if m.IsConsumable() || m.IsSetup() || m.IsEtc() {
			decorator = p.DecorateStackable(compartmentId)
		} else if m.IsCash() || m.IsPet() {
			decorator = p.DecorateCash
		}
		if decorator == nil {
			return Model[any]{}, errors.New("no decorators for reference type")
		}
		return decorator(m)
	}
}

func (p *Processor) GetBySlot(compartmentId uuid.UUID, slot int16) (Model[any], error) {
	return p.BySlotProvider(compartmentId)(slot)()
}

func (p *Processor) BySlotProvider(compartmentId uuid.UUID) func(slot int16) model.Provider[Model[any]] {
	return func(slot int16) model.Provider[Model[any]] {
		t := tenant.MustFromContext(p.ctx)
		ap := model.Map(Make)(getBySlot(t.Id(), compartmentId, slot)(p.db))
		return model.Map(p.DecorateAsset(compartmentId))(ap)
	}
}

func (p *Processor) DecorateEquipable(m Model[any]) (Model[any], error) {
	e, err := equipable.GetById(p.l)(p.ctx)(m.ReferenceId())
	if err != nil {
		return Model[any]{}, err
	}
	return Clone(m).
		SetReferenceData(MakeEquipableReferenceData(e)).
		Build(), nil
}

func MakeEquipableReferenceData(e equipable.Model) EquipableReferenceData {
	return EquipableReferenceData{
		strength:       e.Strength(),
		dexterity:      e.Dexterity(),
		intelligence:   e.Intelligence(),
		luck:           e.Luck(),
		hp:             e.HP(),
		mp:             e.MP(),
		weaponAttack:   e.WeaponAttack(),
		magicAttack:    e.MagicAttack(),
		weaponDefense:  e.WeaponDefense(),
		magicDefense:   e.MagicDefense(),
		accuracy:       e.Accuracy(),
		avoidability:   e.Avoidability(),
		hands:          e.Hands(),
		speed:          e.Speed(),
		jump:           e.Jump(),
		slots:          e.Slots(),
		ownerId:        e.OwnerId(),
		locked:         e.Locked(),
		spikes:         e.Spikes(),
		karmaUsed:      e.KarmaUsed(),
		cold:           e.Cold(),
		canBeTraded:    e.CanBeTraded(),
		levelType:      e.LevelType(),
		level:          e.Level(),
		experience:     e.Experience(),
		hammersApplied: e.HammersApplied(),
		expiration:     e.Expiration(),
	}
}

func (p *Processor) DecorateStackable(compartmentId uuid.UUID) func(m Model[any]) (Model[any], error) {
	return func(m Model[any]) (Model[any], error) {
		s, err := p.stackableProcessor.GetById(m.ReferenceId())
		if err != nil {
			return m, errors.New("cannot locate reference")
		}

		var rd any
		if m.ReferenceType() == ReferenceTypeConsumable {
			rd = ConsumableReferenceData{
				quantity:     s.Quantity(),
				ownerId:      s.OwnerId(),
				flag:         s.Flag(),
				rechargeable: s.Rechargeable(),
			}
		} else if m.ReferenceType() == ReferenceTypeSetup {
			rd = SetupReferenceData{
				quantity: s.Quantity(),
				ownerId:  s.OwnerId(),
				flag:     s.Flag(),
			}
		} else if m.ReferenceType() == ReferenceTypeEtc {
			rd = EtcReferenceData{
				quantity: s.Quantity(),
				ownerId:  s.OwnerId(),
				flag:     s.Flag(),
			}
		}

		return Clone(m).
			SetReferenceData(rd).
			Build(), nil
	}
}

func (p *Processor) DecorateCash(m Model[any]) (Model[any], error) {
	if m.ReferenceType() == ReferenceTypeCash {
		ci, err := p.cashProcessor.GetById(m.ReferenceId())
		if err != nil {
			return m, errors.New("cannot locate reference")
		}
		return Clone(m).
			SetReferenceData(CashReferenceData{
				quantity:   ci.Quantity(),
				ownerId:    ci.OwnerId(),
				flag:       ci.Flag(),
				purchaseBy: ci.PurchasedBy(),
			}).
			Build(), nil
	} else if m.ReferenceType() == ReferenceTypePet {
		ci, err := p.cashProcessor.GetById(m.ReferenceId())
		if err != nil {
			return m, errors.New("cannot locate reference")
		}
		pi, err := pet.GetById(p.l)(p.ctx)(m.ReferenceId())
		if err != nil {
			return m, errors.New("cannot locate reference")
		}
		return Clone(m).
			SetReferenceData(PetReferenceData{
				cashId:     ci.CashId(),
				ownerId:    ci.OwnerId(),
				flag:       ci.Flag(),
				purchaseBy: ci.PurchasedBy(),
				name:       pi.Name(),
				level:      pi.Level(),
				closeness:  pi.Closeness(),
				fullness:   pi.Fullness(),
				expiration: pi.Expiration(),
				slot:       pi.Slot(),
			}).
			Build(), nil
	}
	return m, nil
}

func (p *Processor) Delete(mb *message.Buffer) func(characterId uint32, compartmentId uuid.UUID) func(a Model[any]) error {
	return func(characterId uint32, compartmentId uuid.UUID) func(a Model[any]) error {
		return func(a Model[any]) error {
			t := tenant.MustFromContext(p.ctx)
			p.l.Debugf("Attempting to delete asset [%d].", a.Id())
			txErr := p.db.Transaction(func(tx *gorm.DB) error {
				var deleteRefFunc func(id uint32) error
				if a.ReferenceType() == ReferenceTypeEquipable {
					deleteRefFunc = equipable.Delete(p.l)(p.ctx)
				} else if a.ReferenceType() == ReferenceTypeConsumable || a.ReferenceType() == ReferenceTypeSetup || a.ReferenceType() == ReferenceTypeEtc {
					deleteRefFunc = p.stackableProcessor.Delete
				} else if a.ReferenceType() == ReferenceTypeCash {
					// TODO
				} else if a.ReferenceType() == ReferenceTypePet {
					// TODO
				}

				if deleteRefFunc == nil {
					p.l.Errorf("Unable to locate delete function for asset [%d]. This will lead to a dangling asset.", a.Id())
					return nil
				}
				err := deleteRefFunc(a.ReferenceId())
				if err != nil {
					p.l.WithError(err).Errorf("Unable to delete asset [%d], due to error deleting reference [%d].", a.Id(), a.ReferenceId())
					return err
				}
				err = deleteById(tx, t.Id(), a.Id())
				if err != nil {
					return err
				}
				return mb.Put(asset.EnvEventTopicStatus, asset2.DeletedEventStatusProvider(characterId, compartmentId, a.Id(), a.TemplateId(), a.Slot()))
			})
			if txErr != nil {
				p.l.WithError(txErr).Errorf("Unable to delete asset [%d].", a.Id())
				return txErr
			}
			p.l.Debugf("Deleted asset [%d].", a.Id())
			return nil
		}
	}
}

func (p *Processor) UpdateSlot(mb *message.Buffer) func(characterId uint32, compartmentId uuid.UUID, ap model.Provider[Model[any]], sp model.Provider[int16]) error {
	t := tenant.MustFromContext(p.ctx)
	return func(characterId uint32, compartmentId uuid.UUID, ap model.Provider[Model[any]], sp model.Provider[int16]) error {
		a, err := ap()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err != nil {
			return nil
		}
		s, err := sp()
		if err != nil {
			return err
		}
		p.l.Debugf("Character [%d] attempting to update slot of asset [%d] to [%d] from [%d].", characterId, a.Id(), s, a.Slot())
		err = updateSlot(p.db, t.Id(), a.Id(), s)
		if err != nil {
			return err
		}
		if a.Slot() != int16(math.MinInt16) && s != int16(math.MinInt16) {
			return mb.Put(asset.EnvEventTopicStatus, asset2.MovedEventStatusProvider(characterId, compartmentId, a.Id(), a.TemplateId(), a.Slot(), s))
		}
		return nil
	}
}

func (p *Processor) UpdateQuantity(mb *message.Buffer) func(characterId uint32, compartmentId uuid.UUID, a Model[any], quantity uint32) error {
	return func(characterId uint32, compartmentId uuid.UUID, a Model[any], quantity uint32) error {
		if !a.HasQuantity() {
			return errors.New("cannot update quantity of non-stackable")
		}
		if a.IsConsumable() || a.IsSetup() || a.IsEtc() {
			err := p.stackableProcessor.UpdateQuantity(a.ReferenceId(), quantity)
			if err != nil {
				return err
			}
			return mb.Put(asset.EnvEventTopicStatus, asset2.QuantityChangedEventStatusProvider(characterId, compartmentId, a.Id(), a.TemplateId(), a.Slot(), quantity))
		} else if a.IsCash() {
			err := p.cashProcessor.UpdateQuantity(a.ReferenceId(), quantity)
			if err != nil {
				return err
			}
			return mb.Put(asset.EnvEventTopicStatus, asset2.QuantityChangedEventStatusProvider(characterId, compartmentId, a.Id(), a.TemplateId(), a.Slot(), quantity))
		}
		return errors.New("unknown ReferenceData which implements HasQuantity")
	}
}

func (p *Processor) Create(mb *message.Buffer) func(characterId uint32, compartmentId uuid.UUID, templateId uint32, slot int16, quantity uint32, expiration time.Time, ownerId uint32, flag uint16, rechargeable uint64) (Model[any], error) {
	t := tenant.MustFromContext(p.ctx)
	return func(characterId uint32, compartmentId uuid.UUID, templateId uint32, slot int16, quantity uint32, expiration time.Time, ownerId uint32, flag uint16, rechargeable uint64) (Model[any], error) {
		p.l.Debugf("Character [%d] attempting to create [%d] item(s) [%d] in slot [%d] of compartment [%s].", characterId, quantity, templateId, slot, compartmentId.String())
		var a Model[any]
		txErr := p.db.Transaction(func(tx *gorm.DB) error {
			var referenceId uint32
			var referenceType ReferenceType
			inventoryType, ok := inventory.TypeFromItemId(templateId)
			if !ok {
				return errors.New("unknown item type")
			}
			if inventoryType == inventory.TypeValueEquip {
				e, err := equipable.Create(p.l)(p.ctx)(templateId)()
				if err != nil {
					return err
				}
				referenceId = e.Id()
				referenceType = ReferenceTypeEquipable
			} else if inventoryType == inventory.TypeValueUse {
				s, err := p.stackableProcessor.WithTransaction(tx).Create(compartmentId, quantity, ownerId, flag, rechargeable)
				if err != nil {
					return err
				}
				referenceId = s.Id()
				referenceType = ReferenceTypeConsumable
			} else if inventoryType == inventory.TypeValueSetup {
				s, err := p.stackableProcessor.WithTransaction(tx).Create(compartmentId, quantity, ownerId, flag, rechargeable)
				if err != nil {
					return err
				}
				referenceId = s.Id()
				referenceType = ReferenceTypeSetup
			} else if inventoryType == inventory.TypeValueETC {
				s, err := p.stackableProcessor.WithTransaction(tx).Create(compartmentId, quantity, ownerId, flag, rechargeable)
				if err != nil {
					return err
				}
				referenceId = s.Id()
				referenceType = ReferenceTypeEtc
			} else if inventoryType == inventory.TypeValueCash {
				// TODO
			}

			if referenceId == 0 {
				return errors.New("unknown item type")
			}

			var err error
			a, err = create(p.db, t.Id(), compartmentId, templateId, slot, expiration, referenceId, referenceType)
			if err != nil {
				return err
			}
			return mb.Put(asset.EnvEventTopicStatus, asset2.CreatedEventStatusProvider(characterId, compartmentId, a.Id(), a.TemplateId(), a.Slot()))
		})
		if txErr != nil {
			return Model[any]{}, txErr
		}
		return a, nil
	}
}
