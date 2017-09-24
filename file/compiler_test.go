package file_test

import (
	. "github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/placeholders"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compiler", func() {

	It("adds metadata to the map object", func() {
		source := "../fixtures/maps/meta_and_map.map"

		metaData, mapData := ImportMap(source)
		entityData := placeholders.NewEntityCollection()

		dungeon := CompileLevel(metaData, mapData, entityData)

		Expect(dungeon.Link).To(Equal("prison"))
		Expect(dungeon.Name).To(Equal("The Jovian Prison"))
		Expect(dungeon.Desc).To(Equal("A gloomy building hidden deep in the Jovian woods."))
	})

	It("converts runes into walls and floors", func() {
		source := "../fixtures/maps/meta_and_map.map"

		metaData, mapData := ImportMap(source)
		entityData := placeholders.NewEntityCollection()

		dungeon := CompileLevel(metaData, mapData, entityData)

		Expect(dungeon.Width).To(Equal(mapData.Width))
		Expect(dungeon.Height).To(Equal(mapData.Height))

		// NB: cols/rows are inverted
		walkable := [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		}

		for i := 0; i < len(walkable); i++ {
			for j := 0; j < len(walkable[i]); j++ {
				Expect(dungeon.Tiles[i][j].Walkable).To(Equal(walkable[i][j]))
			}
		}
	})

	It("adds entities to the map", func() {
		map_source := "../fixtures/maps/meta_and_map.map"
		ent_source := "../fixtures/makelists/basic.yaml"

		metaData, mapData := ImportMap(map_source)
		entities := ImportEntities(ent_source)

		dungeon := CompileLevel(metaData, mapData, entities)

		Expect(len(dungeon.Doors)).To(Equal(1))
		Expect(len(dungeon.Mobs)).To(Equal(1))
		Expect(len(dungeon.Keys)).To(Equal(1))
	})

	It("sets the spawn values for each cell", func() {
		map_source := "../fixtures/maps/meta_and_annotated_map.map"
		ent_source := "../fixtures/makelists/annotated.yaml"

		metaData, mapData := ImportMap(map_source)
		entities := ImportEntities(ent_source)

		dungeon := CompileLevel(metaData, mapData, entities)

		Expect(dungeon.Tiles[2][1].Spawn).To(Equal("mob_link"))
	})

	It("Adds the start position", func() {
		map_source := "../fixtures/maps/meta_and_annotated_map.map"
		ent_source := "../fixtures/makelists/annotated.yaml"

		metaData, mapData := ImportMap(map_source)
		entities := ImportEntities(ent_source)

		dungeon := CompileLevel(metaData, mapData, entities)

		Expect(dungeon.StartPosition).To(Equal(placeholders.Position{4, 3}))
	})
})
