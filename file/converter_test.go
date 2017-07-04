package file_test

import (
	. "github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/placeholders"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Converter", func() {

	It("adds metadata to the map object", func() {
		source := "../fixtures/maps/meta_and_map.map"

		metaData, levelData := ImportPlaceholders(source)
		entityData := placeholders.NewEntityCollection()

		level := ConvertPlaceholdersToMap(metaData, levelData, entityData)

		Expect(level.Link).To(Equal("prison"))
		Expect(level.Name).To(Equal("The Jovian Prison"))
		Expect(level.Desc).To(Equal("A gloomy building hidden deep in the Jovian woods."))
	})

	It("converts runes into walls and floors", func() {
		source := "../fixtures/maps/meta_and_map.map"

		metaData, mapData := ImportPlaceholders(source)
		entityData := placeholders.NewEntityCollection()

		level := ConvertPlaceholdersToMap(metaData, mapData, entityData)

		Expect(level.Width).To(Equal(mapData.Width))
		Expect(level.Height).To(Equal(mapData.Height))

		// NB: cols/rows are inverted
		walkable := [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		}

		for i := 0; i < len(walkable); i++ {
			for j := 0; j < len(walkable[i]); j++ {
				Expect(level.Tiles[i][j].Walkable).To(Equal(walkable[i][j]))
			}
		}
	})

	It("adds entities to the map", func() {
		map_source := "../fixtures/maps/meta_and_map.map"
		ent_source := "../fixtures/makelists/basic.yaml"

		metaData, mapData := ImportPlaceholders(map_source)
		entities := ImportEntitiesFromMakeList(ent_source)

		level := ConvertPlaceholdersToMap(metaData, mapData, entities)

		Expect(len(level.Doors)).To(Equal(1))
		Expect(len(level.Mobs)).To(Equal(1))
		Expect(len(level.Keys)).To(Equal(1))
	})

	It("sets the spawn values for each cell", func() {
		map_source := "../fixtures/maps/meta_and_annotated_map.map"
		ent_source := "../fixtures/makelists/annotated.yaml"

		metaData, mapData := ImportPlaceholders(map_source)
		entities := ImportEntitiesFromMakeList(ent_source)

		level := ConvertPlaceholdersToMap(metaData, mapData, entities)

		Expect(level.Tiles[2][1].Spawn).To(Equal("mob_link"))
	})
})
