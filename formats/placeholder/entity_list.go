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

// Mob returns a mob of the same name
func (e *EntityList) Mob(name string) *Mob {
	if mob, exists := e.Mobs[name]; exists {
		return mob
	}
	return nil
}

// Door returns a door of the same name
func (e *EntityList) Door(name string) *Door {
	if door, exists := e.Doors[name]; exists {
		return door
	}
	return nil
}

// Item returns an item of the same name
func (e *EntityList) Item(name string) *Item {
	if item, exists := e.Items[name]; exists {
		return item
	}
	return nil
}
