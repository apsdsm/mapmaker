package output

// Loot contains information about an object held by a mob
type Loot struct {
	Link    string
	MinHeld int
	MaxHeld int
}

// NewLoot returns a new loot object
func NewLoot(link string, min, max int) Loot {
	return Loot{
		Link:    link,
		MinHeld: min,
		MaxHeld: max,
	}
}
