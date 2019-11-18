package file

import "github.com/apsdsm/mapmaker/formats/placeholder"

type DungeonImporter struct {
	Dungeon *placeholder.Dungeon
	Errors  *ErrorList
}

type DungeonImporterConfig struct {
	Errors *ErrorList
}

func NewDungeonImporter(c DungeonImporterConfig) *DungeonImporter {
	i := DungeonImporter{
		Dungeon: &placeholder.Dungeon{},
		Errors:  c.Errors,
	}

	return &i
}

func (i *DungeonImporter) Read(in string) error {

	return nil
}
