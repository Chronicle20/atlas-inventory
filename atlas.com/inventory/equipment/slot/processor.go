package slot

import (
	model2 "atlas-inventory/model"
	"context"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/Chronicle20/atlas-rest/requests"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	l       logrus.FieldLogger
	ctx     context.Context
	GetById func(id uint32) ([]Model, error)
}

func NewProcessor(l logrus.FieldLogger, ctx context.Context) *Processor {
	p := &Processor{
		l:   l,
		ctx: ctx,
	}
	p.GetById = model2.CollapseProvider(p.ByIdModelProvider)
	return p
}

func (p *Processor) ByIdModelProvider(id uint32) model.Provider[[]Model] {
	return requests.SliceProvider[RestModel, Model](p.l, p.ctx)(requestEquipmentSlotDestination(id), Extract, model.Filters[Model]())
}
