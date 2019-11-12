package placeholder

// Door is a placeholder for door entities in a level
type Door struct {
	Reference string            `yaml:"-"`
	Locked    bool              `yaml:"locked"`
	Key       string            `yaml:"key"`
	Events    map[string]string `yaml:"events"`
}

// NeedsKey returns true if the door requires a key
func (d *Door) NeedsKey() bool {
	return d.Key != ""
}
