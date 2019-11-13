package placeholder

// An Item is a single object in the world. Items can be picked up, dropped, or consumed
// according to how they are scripted.
type Item struct {
	Reference string `yaml:"-"`
	Type      string `yaml:"type"`
	Name      string `yaml:"name"`
	Desc      string `yaml:"desc"`
	Size      string `yaml:"size"`
	Uniq      string `yaml:"uniq"`
}
