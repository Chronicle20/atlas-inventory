package asset

import (
	"context"

	"github.com/Chronicle20/atlas-model/model"
	tenant "github.com/Chronicle20/atlas-tenant"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ByCompartmentIdProvider(ctx context.Context) func(db *gorm.DB) func(compartmentId uuid.UUID) model.Provider[[]Model] {
	t := tenant.MustFromContext(ctx)
	return func(db *gorm.DB) func(compartmentId uuid.UUID) model.Provider[[]Model] {
		return func(compartmentId uuid.UUID) model.Provider[[]Model] {
			return model.SliceMap(Make)(getByCompartmentId(t.Id(), compartmentId)(db))(model.ParallelMap())
		}
	}
}

func GetByCompartmentId(l logrus.FieldLogger) func(ctx context.Context) func(db *gorm.DB) func(compartmentId uuid.UUID) ([]Model, error) {
	return func(ctx context.Context) func(db *gorm.DB) func(compartmentId uuid.UUID) ([]Model, error) {
		return func(db *gorm.DB) func(compartmentId uuid.UUID) ([]Model, error) {
			return func(compartmentId uuid.UUID) ([]Model, error) {
				return ByCompartmentIdProvider(ctx)(db)(compartmentId)()
			}
		}
	}
}
