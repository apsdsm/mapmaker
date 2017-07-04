package file

import (
	"github.com/apsdsm/mapmaker/maps"
	"github.com/apsdsm/mapmaker/placeholders"
)

// CompileLevel converts placeholder information to a map file that can be saved as json
func CompileLevel(metaData *placeholders.Meta, mapData *placeholders.Map, entities *placeholders.EntityCollection) *maps.Map {
	m := maps.NewMap(mapData.Width, mapData.Height)

	for x := 0; x < mapData.Width; x++ {
		for y := 0; y < mapData.Height; y++ {
			m.Tiles[x][y].Rune = mapData.Grid[x][y].Rune

			if mapData.Grid[x][y].Rune == ' ' {
				m.Tiles[x][y].Walkable = true
			}
			if mapData.Grid[x][y].Type == "mob" {
				m.Tiles[x][y].Spawn = mapData.Grid[x][y].Link
			}
		}
	}

	m.Link = metaData.Link
	m.Name = metaData.Name
	m.Desc = metaData.Desc
	m.Doors = entities.Doors
	m.Mobs = entities.Mobs
	m.Keys = entities.Keys

	return m
}
