package pet

import (
	"context"

	"github.com/Chronicle20/atlas-model/model"
	"github.com/Chronicle20/atlas-rest/requests"
	"github.com/sirupsen/logrus"
)

func ByIdProvider(l logrus.FieldLogger) func(ctx context.Context) func(petId uint32) model.Provider[Model] {
	return func(ctx context.Context) func(petId uint32) model.Provider[Model] {
		return func(petId uint32) model.Provider[Model] {
			return requests.Provider[RestModel, Model](l, ctx)(requestById(petId), Extract)
		}
	}
}

func GetById(l logrus.FieldLogger) func(ctx context.Context) func(petId uint32) (Model, error) {
	return func(ctx context.Context) func(petId uint32) (Model, error) {
		return func(petId uint32) (Model, error) {
			return ByIdProvider(l)(ctx)(petId)()
		}
	}
}
