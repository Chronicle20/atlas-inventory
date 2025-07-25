package compartment

import (
	"atlas-inventory/compartment"
	consumer2 "atlas-inventory/kafka/consumer"
	compartment2 "atlas-inventory/kafka/message/compartment"
	"context"
	"github.com/Chronicle20/atlas-constants/channel"
	"github.com/Chronicle20/atlas-constants/inventory"
	_map "github.com/Chronicle20/atlas-constants/map"
	"github.com/Chronicle20/atlas-constants/world"
	"github.com/Chronicle20/atlas-kafka/consumer"
	"github.com/Chronicle20/atlas-kafka/handler"
	"github.com/Chronicle20/atlas-kafka/message"
	"github.com/Chronicle20/atlas-kafka/topic"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
)

func InitConsumers(l logrus.FieldLogger) func(func(config consumer.Config, decorators ...model.Decorator[consumer.Config])) func(consumerGroupId string) {
	return func(rf func(config consumer.Config, decorators ...model.Decorator[consumer.Config])) func(consumerGroupId string) {
		return func(consumerGroupId string) {
			rf(consumer2.NewConfig(l)("compartment_command")(compartment2.EnvCommandTopic)(consumerGroupId), consumer.SetHeaderParsers(consumer.SpanHeaderParser, consumer.TenantHeaderParser))
		}
	}
}

func InitHandlers(l logrus.FieldLogger) func(db *gorm.DB) func(rf func(topic string, handler handler.Handler) (string, error)) {
	return func(db *gorm.DB) func(rf func(topic string, handler handler.Handler) (string, error)) {
		return func(rf func(topic string, handler handler.Handler) (string, error)) {
			var t string
			t, _ = topic.EnvProvider(l)(compartment2.EnvCommandTopic)()
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleEquipItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleUnequipItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleMoveItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleDropItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleRequestReserveItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleConsumeItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleDestroyItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleCancelItemReservationCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleIncreaseCapacityCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleCreateAssetCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleRechargeItemCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleMergeCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleSortCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleAcceptCommand(db))))
			_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleReleaseCommand(db))))
		}
	}
}

func handleEquipItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.EquipCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.EquipCommandBody]) {
		if c.Type != compartment2.CommandEquip {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).EquipItemAndEmit(c.TransactionId, c.CharacterId, c.Body.Source, c.Body.Destination)
	}
}

func handleUnequipItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.UnequipCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.UnequipCommandBody]) {
		if c.Type != compartment2.CommandUnequip {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).RemoveEquipAndEmit(c.TransactionId, c.CharacterId, c.Body.Source, c.Body.Destination)
	}
}

func handleMoveItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.MoveCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.MoveCommandBody]) {
		if c.Type != compartment2.CommandMove {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).MoveAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.Source, c.Body.Destination)
	}
}

func handleIncreaseCapacityCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.IncreaseCapacityCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.IncreaseCapacityCommandBody]) {
		if c.Type != compartment2.CommandIncreaseCapacity {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).IncreaseCapacityAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.Amount)
	}
}

func handleDropItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.DropCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.DropCommandBody]) {
		if c.Type != compartment2.CommandDrop {
			return
		}

		m := _map.NewModel(world.Id(c.Body.WorldId))(channel.Id(c.Body.ChannelId))(_map.Id(c.Body.MapId))
		_ = compartment.NewProcessor(l, ctx, db).DropAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType), m, c.Body.X, c.Body.Y, c.Body.Source, c.Body.Quantity)
	}
}

func handleRequestReserveItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.RequestReserveCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.RequestReserveCommandBody]) {
		if c.Type != compartment2.CommandRequestReserve {
			return
		}
		reserves := make([]compartment.ReservationRequest, 0)
		for _, i := range c.Body.Items {
			reserves = append(reserves, compartment.ReservationRequest{
				Slot:     i.Source,
				ItemId:   i.ItemId,
				Quantity: i.Quantity,
			})
		}

		// TODO producers of this command need to be updated to use main TransactionId and not Body.TransactionId
		transactionId := c.TransactionId
		if transactionId == uuid.Nil {
			transactionId = c.Body.TransactionId
		}
		_ = compartment.NewProcessor(l, ctx, db).RequestReserveAndEmit(transactionId, c.CharacterId, inventory.Type(c.InventoryType), reserves)
	}
}

func handleCancelItemReservationCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.CancelReservationCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.CancelReservationCommandBody]) {
		if c.Type != compartment2.CommandCancelReservation {
			return
		}

		// TODO producers of this command need to be updated to use main TransactionId and not Body.TransactionId
		transactionId := c.TransactionId
		if transactionId == uuid.Nil {
			transactionId = c.Body.TransactionId
		}
		_ = compartment.NewProcessor(l, ctx, db).CancelReservationAndEmit(transactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.Slot)
	}
}

func handleConsumeItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.ConsumeCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.ConsumeCommandBody]) {
		if c.Type != compartment2.CommandConsume {
			return
		}

		// TODO producers of this command need to be updated to use main TransactionId and not Body.TransactionId
		transactionId := c.TransactionId
		if transactionId == uuid.Nil {
			transactionId = c.Body.TransactionId
		}
		_ = compartment.NewProcessor(l, ctx, db).ConsumeAssetAndEmit(transactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.Slot)
	}
}

func handleDestroyItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.DestroyCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.DestroyCommandBody]) {
		if c.Type != compartment2.CommandDestroy {
			return
		}
		quantity := c.Body.Quantity
		if quantity == 0 {
			quantity = math.MaxInt32
		}
		_ = compartment.NewProcessor(l, ctx, db).DestroyAssetAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.Slot, quantity)
	}
}

func handleCreateAssetCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.CreateAssetCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.CreateAssetCommandBody]) {
		if c.Type != compartment2.CommandCreateAsset {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).CreateAssetAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.TemplateId, c.Body.Quantity, c.Body.Expiration, c.Body.OwnerId, c.Body.Flag, c.Body.Rechargeable)
	}
}

func handleRechargeItemCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.RechargeCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.RechargeCommandBody]) {
		if c.Type != compartment2.CommandRecharge {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).RechargeAssetAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.Slot, c.Body.Quantity)
	}
}

func handleMergeCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.MergeCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.MergeCommandBody]) {
		if c.Type != compartment2.CommandMerge {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).MergeAndCompactAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType))
	}
}

func handleSortCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.SortCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.SortCommandBody]) {
		if c.Type != compartment2.CommandSort {
			return
		}
		_ = compartment.NewProcessor(l, ctx, db).CompactAndSortAndEmit(c.TransactionId, c.CharacterId, inventory.Type(c.InventoryType))
	}
}

func handleAcceptCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.AcceptCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.AcceptCommandBody]) {
		if c.Type != compartment2.CommandAccept {
			return
		}

		// TODO producers of this command need to be updated to use main TransactionId and not Body.TransactionId
		transactionId := c.TransactionId
		if transactionId == uuid.Nil {
			transactionId = c.Body.TransactionId
		}
		_ = compartment.NewProcessor(l, ctx, db).AcceptAndEmit(transactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.ReferenceId)
	}
}

func handleReleaseCommand(db *gorm.DB) message.Handler[compartment2.Command[compartment2.ReleaseCommandBody]] {
	return func(l logrus.FieldLogger, ctx context.Context, c compartment2.Command[compartment2.ReleaseCommandBody]) {
		if c.Type != compartment2.CommandRelease {
			return
		}

		// TODO producers of this command need to be updated to use main TransactionId and not Body.TransactionId
		transactionId := c.TransactionId
		if transactionId == uuid.Nil {
			transactionId = c.Body.TransactionId
		}
		_ = compartment.NewProcessor(l, ctx, db).ReleaseAndEmit(transactionId, c.CharacterId, inventory.Type(c.InventoryType), c.Body.AssetId)
	}
}
