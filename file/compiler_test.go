// Copyright 2017 Nick del Pozo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file_test

import (
	. "github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/placeholders"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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
			{false, false, false}, {false, true, false}, {false, false, false},
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

	// 0 1 2
	// 3 4 5
	// 6 7 8
	DescribeTable("it calculates tile neighbors",
		func(check int, neighbors [8]int) {

			map_source := "../fixtures/maps/meta_and_map.map"
			ent_source := "../fixtures/makelists/annotated.yaml"

			metaData, mapData := ImportMap(map_source)
			entities := ImportEntities(ent_source)

			dungeon := CompileLevel(metaData, mapData, entities)

			// tile coords
			tile := make(map[int]placeholders.Position)
			tile[-1] = placeholders.Position{-1, -1}
			tile[0] = placeholders.Position{0, 0}
			tile[1] = placeholders.Position{1, 0}
			tile[2] = placeholders.Position{2, 0}
			tile[3] = placeholders.Position{0, 1}
			tile[4] = placeholders.Position{1, 1}
			tile[5] = placeholders.Position{2, 1}
			tile[6] = placeholders.Position{0, 2}
			tile[7] = placeholders.Position{1, 2}
			tile[8] = placeholders.Position{2, 2}

			// check this tile
			c := dungeon.Tiles[tile[check].X][tile[check].Y]

			for i, n := range neighbors {
				Expect(c.Neighbors[i]).To(Equal(tile[n]))
			}
		},
		Entry("tile 0", 0, [8]int{-1, -1, 1, 4, 3, -1, -1, -1}),
		Entry("tile 1", 1, [8]int{-1, -1, 2, 5, 4, 3, 0, -1}),
		Entry("tile 2", 2, [8]int{-1, -1, -1, -1, 5, 4, 1, -1}),
		Entry("tile 3", 3, [8]int{0, 1, 4, 7, 6, -1, -1, -1}),
		Entry("tile 4", 4, [8]int{1, 2, 5, 8, 7, 6, 3, 0}),
		Entry("tile 5", 5, [8]int{2, -1, -1, -1, 8, 7, 4, 1}),
		Entry("tile 6", 6, [8]int{3, 4, 7, -1, -1, -1, -1, -1}),
		Entry("tile 7", 7, [8]int{4, 5, 8, -1, -1, -1, 6, 3}),
		Entry("tile 8", 8, [8]int{5, -1, -1, -1, -1, -1, 7, 4}),
	)
})
