// Copyright 2017 Nick del Pozo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package placeholder

// An EntityCollection is a placeholder collection of the entities created for a level
type EntityCollection struct {
	Mobs  []Mob
	Doors []Door
	Keys  []Key
	Items []Item
}

// NewEntityCollection initializes and returns an EntityCollection object
func NewEntityCollection() EntityCollection {
	e := EntityCollection{
		Mobs:  make([]Mob, 0, 10),
		Doors: make([]Door, 0, 10),
		Keys:  make([]Key, 0, 10),
		Items: make([]Item, 0, 10),
	}

	return e
}

// Mob returns a mob if it exists in the collection. Otherwise nil
func (collection *EntityCollection) Mob(link string) *Mob {
	for _, m := range collection.Mobs {
		if m.Reference == link {
			return &m
		}
	}

	return nil
}

// HasMob returns true if this collection has a mob that is referenced by the given link
func (collection *EntityCollection) HasMob(link string) bool {
	return collection.Mob(link) != nil
}

// AddMob adds mobs to the collection
func (collection *EntityCollection) AddMobs(mob ...Mob) {
	collection.Mobs = append(collection.Mobs, mob...)
}

// GetDoor returns a link to a door if it exists, otherwise nil
func (collection *EntityCollection) GetDoor(link string) *Door {
	for _, d := range collection.Doors {
		if d.Link == link {
			return &d
		}
	}

	return nil
}

// AddDoors adds mobs to the collection
func (collection *EntityCollection) AddDoors(door ...Door) {
	collection.Doors = append(collection.Doors, door...)
}

// HasDoor returns true if this collection has a door that is referenced by the given link
func (collection *EntityCollection) HasDoor(link string) bool {
	return collection.GetDoor(link) != nil
}

// HasKey returns true if key with matching link exists
func (c *EntityCollection) HasKey(link string) bool {
	return c.Key(link) != nil
}

// Key returns address of a key with matching link, or nil if no key exists
func (collection *EntityCollection) Key(link string) *Key {
	for i := range collection.Keys {
		if collection.Keys[i].Link == link {
			return &collection.Keys[i]
		}
	}

	return nil
}

// AddItems adds items to the collection
func (collection *EntityCollection) AddItems(item ...Item) {
	collection.Items = append(collection.Items, item...)
}

// HasItem returns true if this collection has an item that is referenced by the given link
func (collection *EntityCollection) HasItem(link string) bool {
	return collection.Item(link) != nil
}

// Item returns the address of an item with matching link, or nil if no item exists
func (c *EntityCollection) Item(link string) *Item {
	for i := range c.Items {
		if c.Items[i].Link == link {
			return &c.Items[i]
		}
	}
	return nil
}
