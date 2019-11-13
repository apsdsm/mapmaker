package placeholder

// A Door is a baricade that can be either open or closed. When closed it is not walkable.
// Opening the baricade requires using a Key.
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
