package placeholder

type Dungeon struct {
	Link          string    `yaml:"link"`
	Name          string    `yaml:"name"`
	Description   string    `yaml:"desc"`
	Height        int       `yaml:"-"`
	Width         int       `yaml:"-"`
	Grid          [][]*Cell `yaml:"-"`
	StartPosition *Position `yaml:"-"`
}

// AllocateTiles will generate enough space for the tiles requires by width and height
func (d *Dungeon) AllocateTiles(width, height int) {
	d.Width = width
	d.Height = height
	d.Grid = make([][]*Cell, width)

	for i := 0; i < width; i++ {
		d.Grid[i] = make([]*Cell, height)
	}
}

// NewDungeon returns an initialized dungeon placeholder
func NewDungeon(height, width int) *Dungeon {
	m := Dungeon{
		Width:  width,
		Height: height,
		Grid:   make([][]*Cell, width),
	}

	for i := 0; i < width; i++ {
		m.Grid[i] = make([]*Cell, height)
	}

	return &m
}
