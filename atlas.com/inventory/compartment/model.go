package compartment

import (
	"atlas-inventory/asset"
	"errors"
	"sort"

	"github.com/Chronicle20/atlas-constants/inventory"
	"github.com/google/uuid"
)

type Model struct {
	id            uuid.UUID
	characterId   uint32
	inventoryType inventory.Type
	capacity      uint32
	assets        []asset.Model[any]
}

func (m Model) Id() uuid.UUID {
	return m.id
}

func (m Model) Type() inventory.Type {
	return m.inventoryType
}

func (m Model) Capacity() uint32 {
	return m.capacity
}

func (m Model) Assets() []asset.Model[any] {
	return m.assets
}

func (m Model) CharacterId() uint32 {
	return m.characterId
}

func (m Model) NextFreeSlot() (int16, error) {
	if len(m.Assets()) == 0 {
		return 1, nil
	}
	sort.Slice(m.Assets(), func(i, j int) bool {
		return m.Assets()[i].Slot() < m.Assets()[j].Slot()
	})

	slot := int16(1)
	i := 0

	for {
		if slot > int16(m.Capacity()) {
			return 0, errors.New("no free slots")
		} else if i >= len(m.Assets()) {
			return slot, nil
		} else if slot < m.Assets()[i].Slot() {
			return slot, nil
		} else if slot == m.Assets()[i].Slot() {
			slot += 1
			i += 1
		} else if m.Assets()[i].Slot() <= 0 {
			i += 1
		}
	}
}

func Clone(m Model) *ModelBuilder {
	return &ModelBuilder{
		id:            m.id,
		characterId:   m.characterId,
		inventoryType: m.inventoryType,
		capacity:      m.capacity,
		assets:        m.assets,
	}
}

type ModelBuilder struct {
	id            uuid.UUID
	characterId   uint32
	inventoryType inventory.Type
	capacity      uint32
	assets        []asset.Model[any]
}

func NewBuilder(id uuid.UUID, characterId uint32, it inventory.Type, capacity uint32) *ModelBuilder {
	return &ModelBuilder{
		id:            id,
		characterId:   characterId,
		inventoryType: it,
		capacity:      capacity,
		assets:        make([]asset.Model[any], 0),
	}
}

func (b *ModelBuilder) SetCapacity(capacity uint32) *ModelBuilder {
	b.capacity = capacity
	return b
}

func (b *ModelBuilder) AddAsset(a asset.Model[any]) *ModelBuilder {
	b.assets = append(b.assets, a)
	return b
}

func (b *ModelBuilder) SetAssets(as []asset.Model[any]) *ModelBuilder {
	b.assets = as
	return b
}

func (b *ModelBuilder) Build() Model {
	return Model{
		id:            b.id,
		characterId:   b.characterId,
		inventoryType: b.inventoryType,
		capacity:      b.capacity,
		assets:        b.assets,
	}
}
