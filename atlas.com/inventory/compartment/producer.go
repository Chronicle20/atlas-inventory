package compartment

import (
	"atlas-inventory/kafka/message/compartment"
	"github.com/Chronicle20/atlas-constants/inventory"
	"github.com/Chronicle20/atlas-kafka/producer"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func CreatedEventStatusProvider(id uuid.UUID, characterId uint32, inventoryType inventory.Type, capacity uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.CreatedStatusEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeCreated,
		Body: compartment.CreatedStatusEventBody{
			Type:     byte(inventoryType),
			Capacity: capacity,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func DeletedEventStatusProvider(id uuid.UUID, characterId uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.DeletedStatusEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeDeleted,
		Body:          compartment.DeletedStatusEventBody{},
	}
	return producer.SingleMessageProvider(key, value)
}

func CapacityChangedEventStatusProvider(id uuid.UUID, characterId uint32, inventoryType inventory.Type, capacity uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.CapacityChangedEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeCapacityChanged,
		Body: compartment.CapacityChangedEventBody{
			Type:     byte(inventoryType),
			Capacity: capacity,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func ReservedEventStatusProvider(id uuid.UUID, characterId uint32, itemId uint32, slot int16, quantity uint32, transactionId uuid.UUID) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.ReservedEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeReserved,
		Body: compartment.ReservedEventBody{
			TransactionId: transactionId,
			ItemId:        itemId,
			Slot:          slot,
			Quantity:      quantity,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func ReservationCancelledEventStatusProvider(id uuid.UUID, characterId uint32, itemId uint32, slot int16, quantity uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.ReservationCancelledEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeReservationCancelled,
		Body: compartment.ReservationCancelledEventBody{
			ItemId:   itemId,
			Slot:     slot,
			Quantity: quantity,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func MergeCompleteEventStatusProvider(id uuid.UUID, characterId uint32, inventoryType inventory.Type) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.MergeCompleteEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeMergeComplete,
		Body: compartment.MergeCompleteEventBody{
			Type: byte(inventoryType),
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func SortCompleteEventStatusProvider(id uuid.UUID, characterId uint32, inventoryType inventory.Type) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &compartment.StatusEvent[compartment.SortCompleteEventBody]{
		CharacterId:   characterId,
		CompartmentId: id,
		Type:          compartment.StatusEventTypeSortComplete,
		Body: compartment.SortCompleteEventBody{
			Type: byte(inventoryType),
		},
	}
	return producer.SingleMessageProvider(key, value)
}
