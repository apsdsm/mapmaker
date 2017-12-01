package placeholder

// Loot is a placeholder structure for imported mob items
type Loot struct {
	Raw     string
	Link    string
	MaxHeld int
	MinHeld int
}

// NewLoot returns a new loot object
func NewLoot(raw, link string, min, max int) Loot {
	return Loot{
		Raw:     raw,
		Link:    link,
		MinHeld: min,
		MaxHeld: max,
	}
}
