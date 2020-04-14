package placeholder

// EntityList is a representation of a list of importable entities. It should be imported
// using the EntityImporter.
type EntityList struct {
	Mobs  map[string]*Mob  `yaml:"mobs"`
	Doors map[string]*Door `yaml:"doors"`
	Items map[string]*Item `yaml:"items"`
}

// NewEntityList initializes and returns a new EntityList.
func NewEntityList() *EntityList {
	return &EntityList{
		Mobs:  make(map[string]*Mob),
		Doors: make(map[string]*Door),
		Items: make(map[string]*Item),
	}
}

// HasMob returns true if the mob is defined in the entity list
func (e *EntityList) HasMob(name string) bool {
	_, exists := e.Mobs[name]
	return exists
}

// Mob returns a mob of the same name
func (e *EntityList) Mob(name string) *Mob {
	if mob, exists := e.Mobs[name]; exists {
		return mob
	}
	return nil
}

// HasDoor returns true if the door is defined in the entity list
func (e *EntityList) HasDoor(name string) bool {
	_, exists := e.Doors[name]
	return exists
}

// Door returns a door of the same name
func (e *EntityList) Door(name string) *Door {
	if door, exists := e.Doors[name]; exists {
		return door
	}
	return nil
}

// HasItem returns true if the item is defined in the entity list
func (e *EntityList) HasItem(name string) bool {
	_, exists := e.Items[name]
	return exists
}

// Item returns an item of the same name
func (e *EntityList) Item(name string) *Item {
	if item, exists := e.Items[name]; exists {
		return item
	}
	return nil
}
