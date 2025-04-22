package asset

import "time"

type EquipableReferenceData struct {
	strength       uint16
	dexterity      uint16
	intelligence   uint16
	luck           uint16
	hp             uint16
	mp             uint16
	weaponAttack   uint16
	magicAttack    uint16
	weaponDefense  uint16
	magicDefense   uint16
	accuracy       uint16
	avoidability   uint16
	hands          uint16
	speed          uint16
	jump           uint16
	slots          uint16
	ownerId        uint32
	locked         bool
	spikes         bool
	karmaUsed      bool
	cold           bool
	canBeTraded    bool
	levelType      byte
	level          byte
	experience     uint32
	hammersApplied uint32
	expiration     time.Time
}

func (e EquipableReferenceData) Strength() uint16       { return e.strength }
func (e EquipableReferenceData) Dexterity() uint16      { return e.dexterity }
func (e EquipableReferenceData) Intelligence() uint16   { return e.intelligence }
func (e EquipableReferenceData) Luck() uint16           { return e.luck }
func (e EquipableReferenceData) HP() uint16             { return e.hp }
func (e EquipableReferenceData) MP() uint16             { return e.mp }
func (e EquipableReferenceData) WeaponAttack() uint16   { return e.weaponAttack }
func (e EquipableReferenceData) MagicAttack() uint16    { return e.magicAttack }
func (e EquipableReferenceData) WeaponDefense() uint16  { return e.weaponDefense }
func (e EquipableReferenceData) MagicDefense() uint16   { return e.magicDefense }
func (e EquipableReferenceData) Accuracy() uint16       { return e.accuracy }
func (e EquipableReferenceData) Avoidability() uint16   { return e.avoidability }
func (e EquipableReferenceData) Hands() uint16          { return e.hands }
func (e EquipableReferenceData) Speed() uint16          { return e.speed }
func (e EquipableReferenceData) Jump() uint16           { return e.jump }
func (e EquipableReferenceData) Slots() uint16          { return e.slots }
func (e EquipableReferenceData) OwnerId() uint32        { return e.ownerId }
func (e EquipableReferenceData) IsLocked() bool         { return e.locked }
func (e EquipableReferenceData) HasSpikes() bool        { return e.spikes }
func (e EquipableReferenceData) IsKarmaUsed() bool      { return e.karmaUsed }
func (e EquipableReferenceData) IsCold() bool           { return e.cold }
func (e EquipableReferenceData) CanBeTraded() bool      { return e.canBeTraded }
func (e EquipableReferenceData) LevelType() byte        { return e.levelType }
func (e EquipableReferenceData) Level() byte            { return e.level }
func (e EquipableReferenceData) Experience() uint32     { return e.experience }
func (e EquipableReferenceData) HammersApplied() uint32 { return e.hammersApplied }
func (e EquipableReferenceData) Expiration() time.Time  { return e.expiration }

type EquipableReferenceDataBuilder struct {
	strength       uint16
	dexterity      uint16
	intelligence   uint16
	luck           uint16
	hp             uint16
	mp             uint16
	weaponAttack   uint16
	magicAttack    uint16
	weaponDefense  uint16
	magicDefense   uint16
	accuracy       uint16
	avoidability   uint16
	hands          uint16
	speed          uint16
	jump           uint16
	slots          uint16
	ownerId        uint32
	locked         bool
	spikes         bool
	karmaUsed      bool
	cold           bool
	canBeTraded    bool
	levelType      byte
	level          byte
	experience     uint32
	hammersApplied uint32
	expiration     time.Time
}

// NewEquipableReferenceDataBuilder creates a new builder instance.
func NewEquipableReferenceDataBuilder() *EquipableReferenceDataBuilder {
	return &EquipableReferenceDataBuilder{}
}

// Clone initializes the builder with data from the provided model.
func (b *EquipableReferenceDataBuilder) Clone(model EquipableReferenceData) *EquipableReferenceDataBuilder {
	*b = EquipableReferenceDataBuilder{
		strength:       model.strength,
		dexterity:      model.dexterity,
		intelligence:   model.intelligence,
		luck:           model.luck,
		hp:             model.hp,
		mp:             model.mp,
		weaponAttack:   model.weaponAttack,
		magicAttack:    model.magicAttack,
		weaponDefense:  model.weaponDefense,
		magicDefense:   model.magicDefense,
		accuracy:       model.accuracy,
		avoidability:   model.avoidability,
		hands:          model.hands,
		speed:          model.speed,
		jump:           model.jump,
		slots:          model.slots,
		ownerId:        model.ownerId,
		locked:         model.locked,
		spikes:         model.spikes,
		karmaUsed:      model.karmaUsed,
		cold:           model.cold,
		canBeTraded:    model.canBeTraded,
		levelType:      model.levelType,
		level:          model.level,
		experience:     model.experience,
		hammersApplied: model.hammersApplied,
		expiration:     model.expiration,
	}
	return b
}

// Build assembles the final EquipableReferenceData from the builder.
func (b *EquipableReferenceDataBuilder) Build() EquipableReferenceData {
	return EquipableReferenceData{
		strength:       b.strength,
		dexterity:      b.dexterity,
		intelligence:   b.intelligence,
		luck:           b.luck,
		hp:             b.hp,
		mp:             b.mp,
		weaponAttack:   b.weaponAttack,
		magicAttack:    b.magicAttack,
		weaponDefense:  b.weaponDefense,
		magicDefense:   b.magicDefense,
		accuracy:       b.accuracy,
		avoidability:   b.avoidability,
		hands:          b.hands,
		speed:          b.speed,
		jump:           b.jump,
		slots:          b.slots,
		ownerId:        b.ownerId,
		locked:         b.locked,
		spikes:         b.spikes,
		karmaUsed:      b.karmaUsed,
		cold:           b.cold,
		canBeTraded:    b.canBeTraded,
		levelType:      b.levelType,
		level:          b.level,
		experience:     b.experience,
		hammersApplied: b.hammersApplied,
		expiration:     b.expiration,
	}
}

func (b *EquipableReferenceDataBuilder) SetStrength(value uint16) *EquipableReferenceDataBuilder {
	b.strength = value
	return b
}

// Setters for EquipableReferenceDataBuilder

func (b *EquipableReferenceDataBuilder) SetDexterity(value uint16) *EquipableReferenceDataBuilder {
	b.dexterity = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetIntelligence(value uint16) *EquipableReferenceDataBuilder {
	b.intelligence = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetLuck(value uint16) *EquipableReferenceDataBuilder {
	b.luck = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetHp(value uint16) *EquipableReferenceDataBuilder {
	b.hp = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetMp(value uint16) *EquipableReferenceDataBuilder {
	b.mp = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetWeaponAttack(value uint16) *EquipableReferenceDataBuilder {
	b.weaponAttack = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetMagicAttack(value uint16) *EquipableReferenceDataBuilder {
	b.magicAttack = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetWeaponDefense(value uint16) *EquipableReferenceDataBuilder {
	b.weaponDefense = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetMagicDefense(value uint16) *EquipableReferenceDataBuilder {
	b.magicDefense = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetAccuracy(value uint16) *EquipableReferenceDataBuilder {
	b.accuracy = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetAvoidability(value uint16) *EquipableReferenceDataBuilder {
	b.avoidability = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetHands(value uint16) *EquipableReferenceDataBuilder {
	b.hands = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetSpeed(value uint16) *EquipableReferenceDataBuilder {
	b.speed = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetJump(value uint16) *EquipableReferenceDataBuilder {
	b.jump = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetSlots(value uint16) *EquipableReferenceDataBuilder {
	b.slots = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetOwnerId(value uint32) *EquipableReferenceDataBuilder {
	b.ownerId = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetLocked(value bool) *EquipableReferenceDataBuilder {
	b.locked = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetSpikes(value bool) *EquipableReferenceDataBuilder {
	b.spikes = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetKarmaUsed(value bool) *EquipableReferenceDataBuilder {
	b.karmaUsed = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetCold(value bool) *EquipableReferenceDataBuilder {
	b.cold = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetCanBeTraded(value bool) *EquipableReferenceDataBuilder {
	b.canBeTraded = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetLevelType(value byte) *EquipableReferenceDataBuilder {
	b.levelType = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetLevel(value byte) *EquipableReferenceDataBuilder {
	b.level = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetExperience(value uint32) *EquipableReferenceDataBuilder {
	b.experience = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetHammersApplied(value uint32) *EquipableReferenceDataBuilder {
	b.hammersApplied = value
	return b
}

func (b *EquipableReferenceDataBuilder) SetExpiration(value time.Time) *EquipableReferenceDataBuilder {
	b.expiration = value
	return b
}

type ConsumableReferenceData struct {
	quantity     uint32
	ownerId      uint32
	flag         uint16
	rechargeable uint64
}

func (c ConsumableReferenceData) Quantity() uint32 {
	return c.quantity
}

func (c ConsumableReferenceData) OwnerId() uint32 {
	return c.ownerId
}

func (c ConsumableReferenceData) Flag() uint16 {
	return c.flag
}

func (c ConsumableReferenceData) Rechargeable() uint64 {
	return c.rechargeable
}

type ConsumableReferenceDataBuilder struct {
	quantity     uint32
	ownerId      uint32
	flag         uint16
	rechargeable uint64
}

func NewConsumableReferenceDataBuilder() *ConsumableReferenceDataBuilder {
	return &ConsumableReferenceDataBuilder{}
}

func (b *ConsumableReferenceDataBuilder) SetQuantity(value uint32) *ConsumableReferenceDataBuilder {
	b.quantity = value
	return b
}

func (b *ConsumableReferenceDataBuilder) SetOwnerId(value uint32) *ConsumableReferenceDataBuilder {
	b.ownerId = value
	return b
}

func (b *ConsumableReferenceDataBuilder) SetFlag(value uint16) *ConsumableReferenceDataBuilder {
	b.flag = value
	return b
}

func (b *ConsumableReferenceDataBuilder) SetRechargeable(value uint64) *ConsumableReferenceDataBuilder {
	b.rechargeable = value
	return b
}

func (b *ConsumableReferenceDataBuilder) Build() ConsumableReferenceData {
	return ConsumableReferenceData{
		quantity:     b.quantity,
		ownerId:      b.ownerId,
		flag:         b.flag,
		rechargeable: b.rechargeable,
	}
}

type SetupReferenceData struct {
	quantity uint32
	ownerId  uint32
	flag     uint16
}

func (c SetupReferenceData) Quantity() uint32 {
	return c.quantity
}

func (c SetupReferenceData) OwnerId() uint32 {
	return c.ownerId
}

func (c SetupReferenceData) Flag() uint16 {
	return c.flag
}

type SetupReferenceDataBuilder struct {
	quantity uint32
	ownerId  uint32
	flag     uint16
}

func NewSetupReferenceDataBuilder() *SetupReferenceDataBuilder {
	return &SetupReferenceDataBuilder{}
}

func (b *SetupReferenceDataBuilder) SetQuantity(value uint32) *SetupReferenceDataBuilder {
	b.quantity = value
	return b
}

func (b *SetupReferenceDataBuilder) SetOwnerId(value uint32) *SetupReferenceDataBuilder {
	b.ownerId = value
	return b
}

func (b *SetupReferenceDataBuilder) SetFlag(value uint16) *SetupReferenceDataBuilder {
	b.flag = value
	return b
}

func (b *SetupReferenceDataBuilder) Build() SetupReferenceData {
	return SetupReferenceData{
		quantity: b.quantity,
		ownerId:  b.ownerId,
		flag:     b.flag,
	}
}

type EtcReferenceData struct {
	quantity uint32
	ownerId  uint32
	flag     uint16
}

func (c EtcReferenceData) Quantity() uint32 {
	return c.quantity
}

func (c EtcReferenceData) OwnerId() uint32 {
	return c.ownerId
}

func (c EtcReferenceData) Flag() uint16 {
	return c.flag
}

type EtcReferenceDataBuilder struct {
	quantity uint32
	ownerId  uint32
	flag     uint16
}

func NewEtcReferenceDataBuilder() *EtcReferenceDataBuilder {
	return &EtcReferenceDataBuilder{}
}

func (b *EtcReferenceDataBuilder) SetQuantity(value uint32) *EtcReferenceDataBuilder {
	b.quantity = value
	return b
}

func (b *EtcReferenceDataBuilder) SetOwnerId(value uint32) *EtcReferenceDataBuilder {
	b.ownerId = value
	return b
}

func (b *EtcReferenceDataBuilder) SetFlag(value uint16) *EtcReferenceDataBuilder {
	b.flag = value
	return b
}

func (b *EtcReferenceDataBuilder) Build() EtcReferenceData {
	return EtcReferenceData{
		quantity: b.quantity,
		ownerId:  b.ownerId,
		flag:     b.flag,
	}
}

type CashReferenceData struct {
	cashId     uint64
	quantity   uint32
	ownerId    uint32
	flag       uint16
	purchaseBy uint32
}

func (c CashReferenceData) Quantity() uint32 {
	return c.quantity
}

func (c CashReferenceData) CashId() uint64 {
	return c.cashId
}

func (c CashReferenceData) OwnerId() uint32 {
	return c.ownerId
}

func (c CashReferenceData) Flag() uint16 {
	return c.flag
}

func (c CashReferenceData) PurchaseBy() uint32 {
	return c.purchaseBy
}

type CashReferenceDataBuilder struct {
	cashId     uint64
	quantity   uint32
	ownerId    uint32
	flag       uint16
	purchaseBy uint32
}

func NewCashReferenceDataBuilder() *CashReferenceDataBuilder {
	return &CashReferenceDataBuilder{}
}

func (b *CashReferenceDataBuilder) SetCashId(value uint64) *CashReferenceDataBuilder {
	b.cashId = value
	return b
}

func (b *CashReferenceDataBuilder) SetQuantity(value uint32) *CashReferenceDataBuilder {
	b.quantity = value
	return b
}

func (b *CashReferenceDataBuilder) SetOwnerId(value uint32) *CashReferenceDataBuilder {
	b.ownerId = value
	return b
}

func (b *CashReferenceDataBuilder) SetFlag(value uint16) *CashReferenceDataBuilder {
	b.flag = value
	return b
}

func (b *CashReferenceDataBuilder) SetPurchaseBy(value uint32) *CashReferenceDataBuilder {
	b.purchaseBy = value
	return b
}

func (b *CashReferenceDataBuilder) Build() CashReferenceData {
	return CashReferenceData{
		cashId:     b.cashId,
		quantity:   b.quantity,
		ownerId:    b.ownerId,
		flag:       b.flag,
		purchaseBy: b.purchaseBy,
	}
}

type PetReferenceData struct {
	cashId        uint64
	ownerId       uint32
	flag          uint16
	purchaseBy    uint32
	name          string
	level         byte
	closeness     uint16
	fullness      byte
	expiration    time.Time
	slot          int8
	attribute     uint16
	skill         uint16
	remainingLife uint32
	attribute2    uint16
}

func (d PetReferenceData) CashId() uint64 {
	return d.cashId
}

func (d PetReferenceData) OwnerId() uint32 {
	return d.ownerId
}

func (d PetReferenceData) Flag() uint16 {
	return d.flag
}

func (d PetReferenceData) PurchaseBy() uint32 {
	return d.purchaseBy
}

func (d PetReferenceData) Name() string {
	return d.name
}

func (d PetReferenceData) Level() byte {
	return d.level
}

func (d PetReferenceData) Closeness() uint16 {
	return d.closeness
}

func (d PetReferenceData) Fullness() byte {
	return d.fullness
}

func (d PetReferenceData) Slot() int8 {
	return d.slot
}

type PetReferenceDataBuilder struct {
	cashId        uint64
	ownerId       uint32
	flag          uint16
	purchaseBy    uint32
	name          string
	level         byte
	closeness     uint16
	fullness      byte
	expiration    time.Time
	slot          int8
	attribute     uint16
	skill         uint16
	remainingLife uint32
	attribute2    uint16
}

func NewPetReferenceDataBuilder() *PetReferenceDataBuilder {
	return &PetReferenceDataBuilder{}
}

func (b *PetReferenceDataBuilder) SetCashId(value uint64) *PetReferenceDataBuilder {
	b.cashId = value
	return b
}

func (b *PetReferenceDataBuilder) SetOwnerId(value uint32) *PetReferenceDataBuilder {
	b.ownerId = value
	return b
}

func (b *PetReferenceDataBuilder) SetFlag(value uint16) *PetReferenceDataBuilder {
	b.flag = value
	return b
}

func (b *PetReferenceDataBuilder) SetPurchaseBy(value uint32) *PetReferenceDataBuilder {
	b.purchaseBy = value
	return b
}

func (b *PetReferenceDataBuilder) SetName(value string) *PetReferenceDataBuilder {
	b.name = value
	return b
}

func (b *PetReferenceDataBuilder) SetLevel(value byte) *PetReferenceDataBuilder {
	b.level = value
	return b
}

func (b *PetReferenceDataBuilder) SetCloseness(value uint16) *PetReferenceDataBuilder {
	b.closeness = value
	return b
}

func (b *PetReferenceDataBuilder) SetFullness(value byte) *PetReferenceDataBuilder {
	b.fullness = value
	return b
}

func (b *PetReferenceDataBuilder) SetExpiration(value time.Time) *PetReferenceDataBuilder {
	b.expiration = value
	return b
}

func (b *PetReferenceDataBuilder) SetSlot(value int8) *PetReferenceDataBuilder {
	b.slot = value
	return b
}

func (b *PetReferenceDataBuilder) SetAttribute(value uint16) *PetReferenceDataBuilder {
	b.attribute = value
	return b
}

func (b *PetReferenceDataBuilder) SetSkill(value uint16) *PetReferenceDataBuilder {
	b.skill = value
	return b
}

func (b *PetReferenceDataBuilder) SetRemainingLife(value uint32) *PetReferenceDataBuilder {
	b.remainingLife = value
	return b
}

func (b *PetReferenceDataBuilder) SetAttribute2(value uint16) *PetReferenceDataBuilder {
	b.attribute2 = value
	return b
}

func (b *PetReferenceDataBuilder) Build() PetReferenceData {
	return PetReferenceData{
		cashId:        b.cashId,
		ownerId:       b.ownerId,
		flag:          b.flag,
		purchaseBy:    b.purchaseBy,
		name:          b.name,
		level:         b.level,
		closeness:     b.closeness,
		fullness:      b.fullness,
		expiration:    b.expiration,
		slot:          b.slot,
		attribute:     b.attribute,
		skill:         b.skill,
		remainingLife: b.remainingLife,
		attribute2:    b.attribute2,
	}
}
