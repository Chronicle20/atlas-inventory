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
