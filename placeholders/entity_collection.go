//    Copyright 2017 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package placeholders

// An EntityCollection is a placeholder collection of the entities created for a level
type EntityCollection struct {
	Mobs  []Mob
	Doors []Door
	Keys  []Key
}

// NewEntityCollection initializes and returns an EntityCollection object
func NewEntityCollection() *EntityCollection {
	e := EntityCollection{
		Mobs:  make([]Mob, 0, 10),
		Doors: make([]Door, 0, 10),
		Keys:  make([]Key, 0, 10),
	}

	return &e
}

// GetMob returns a mob if it exists in the collection. Otherwise nil
func (collection *EntityCollection) GetMob(link string) *Mob {
	for _, m := range collection.Mobs {
		if m.Link == link {
			return &m
		}
	}

	return nil
}

// HasMob returns true if this collection has a mob that is referenced by the given link
func (collection *EntityCollection) HasMob(link string) bool {
	return collection.GetMob(link) != nil
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

// HasDoor returns true if this collection has a door that is referenced by the given link
func (collection *EntityCollection) HasDoor(link string) bool {
	return collection.GetDoor(link) != nil
}

// HasKey returns true if this collection has a key that is referenced by the given link
func (collection *EntityCollection) HasKey(link string) bool {
	for _, k := range collection.Keys {
		if k.Link == link {
			return true
		}
	}

	return false
}
