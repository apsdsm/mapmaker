package placeholder

type EntityList struct {
	Mobs map[string]*Mob `yaml:"mobs"`
}

func NewEntityList() *EntityList {
	return &EntityList{
		Mobs: make(map[string]*Mob, 0),
	}
}

func (e *EntityList) Mob(name string) *Mob {
	if mob, exists := e.Mobs[name]; exists {
		return mob
	}
	return nil
}
