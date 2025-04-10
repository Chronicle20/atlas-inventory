package stackable

import (
	model2 "atlas-inventory/model"
	"context"
	"github.com/Chronicle20/atlas-model/model"
	tenant "github.com/Chronicle20/atlas-tenant"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Processor struct {
	l       logrus.FieldLogger
	ctx     context.Context
	db      *gorm.DB
	GetById func(id uint32) (Model, error)
}

func NewProcessor(l logrus.FieldLogger, ctx context.Context, db *gorm.DB) *Processor {
	p := &Processor{
		l:   l,
		ctx: ctx,
		db:  db,
	}
	p.GetById = model2.CollapseProvider(p.ByIdProvider)
	return p
}

func (p *Processor) WithTransaction(db *gorm.DB) *Processor {
	return &Processor{
		l:       p.l,
		ctx:     p.ctx,
		db:      db,
		GetById: p.GetById,
	}
}

func (p *Processor) ByCompartmentIdProvider(compartmentId uuid.UUID) model.Provider[[]Model] {
	t := tenant.MustFromContext(p.ctx)
	return model.SliceMap(Make)(getByCompartmentId(t.Id(), compartmentId)(p.db))(model.ParallelMap())
}

func (p *Processor) Delete(id uint32) error {
	t := tenant.MustFromContext(p.ctx)
	p.l.Debugf("Attempting to delete stackable item [%d].", id)
	return deleteById(p.db, t.Id(), id)
}

func (p *Processor) UpdateQuantity(id uint32, quantity uint32) error {
	t := tenant.MustFromContext(p.ctx)
	return updateQuantity(p.db, t.Id(), id, quantity)
}

func (p *Processor) ByIdProvider(id uint32) model.Provider[Model] {
	t := tenant.MustFromContext(p.ctx)
	return model.Map(Make)(getById(t.Id(), id)(p.db))
}
