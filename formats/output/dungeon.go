package output

// A Dungeon contains map and entity data for a single dungeon
type Dungeon struct {
	Width, Height    int
	Link, Name, Desc string
	StartPosition    Position
	Tiles            [][]Tile
	Doors            []Door
	Keys             []Key
	Items            []Item
	Mobs             []Mob
}

// NewDungeon generates a new map initialized to the specified size
func NewDungeon(width, height int) *Dungeon {
	m := Dungeon{}

	m.Width = width
	m.Height = height
	m.Tiles = make([][]Tile, width)

	for i := range m.Tiles {
		m.Tiles[i] = make([]Tile, height)
	}

	return &m
}
