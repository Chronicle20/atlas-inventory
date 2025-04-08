package character

import (
	"atlas-inventory/inventory"
	consumer2 "atlas-inventory/kafka/consumer"
	"atlas-inventory/kafka/message/character"
	"context"
	"github.com/Chronicle20/atlas-kafka/consumer"
	"github.com/Chronicle20/atlas-kafka/handler"
	"github.com/Chronicle20/atlas-kafka/message"
	"github.com/Chronicle20/atlas-kafka/topic"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitConsumers(l logrus.FieldLogger) func(func(config consumer.Config, decorators ...model.Decorator[consumer.Config])) func(consumerGroupId string) {
	return func(rf func(config consumer.Config, decorators ...model.Decorator[consumer.Config])) func(consumerGroupId string) {
		return func(consumerGroupId string) {
			rf(consumer2.NewConfig(l)("character_created")(character.EnvEventTopicStatus)(consumerGroupId), consumer.SetHeaderParsers(consumer.SpanHeaderParser, consumer.TenantHeaderParser))
		}
	}
}

func InitHandlers(l logrus.FieldLogger) func(db *gorm.DB) func(rf func(topic string, handler handler.Handler) (string, error)) {
	return func(db *gorm.DB) func(rf func(topic string, handler handler.Handler) (string, error)) {
		return func(rf func(topic string, handler handler.Handler) (string, error)) {
			var t string
			t, _ = topic.EnvProvider(l)(character.EnvEventTopicStatus)()
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleStatusEventCreated(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleStatusEventDeleted(db))))
		}
	}
}

func handleStatusEventCreated(db *gorm.DB) func(l logrus.FieldLogger, ctx context.Context, e character.StatusEvent[character.CreatedStatusBody]) {
	return func(l logrus.FieldLogger, ctx context.Context, e character.StatusEvent[character.CreatedStatusBody]) {
		if e.Type != character.StatusEventTypeCreated {
			return
		}
		_, err := inventory.NewProcessor(l, ctx, db).CreateAndEmit(e.CharacterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to create for character [%d].", e.CharacterId)
		}
	}
}

func handleStatusEventDeleted(db *gorm.DB) message.Handler[character.StatusEvent[character.DeletedStatusEventBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, e character.StatusEvent[character.DeletedStatusEventBody]) {
		if e.Type != character.StatusEventTypeDeleted {
			return
		}
		err := inventory.NewProcessor(l, ctx, db).DeleteAndEmit(e.CharacterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to delete for character [%d].", e.CharacterId)
		}
	}
}
