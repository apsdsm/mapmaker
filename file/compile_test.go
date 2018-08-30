//    Copyright 2017 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package file_test

import (
	. "github.com/apsdsm/mapmaker/file"
	"github.com/apsdsm/mapmaker/formats/output"
	"github.com/apsdsm/mapmaker/formats/placeholder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compiler", func() {

	It("adds metadata to the map object", func() {
		dungeon := placeholder.NewMap(0, 0)

		meta := placeholder.Meta{
			Link: "meta_link",
			Name: "meta_name",
			Desc: "meta_desc",
		}

		entities := placeholder.NewEntityCollection()

		compiled := Compile(&meta, dungeon, &entities)

		Expect(compiled.Link).To(Equal("meta_link"))
		Expect(compiled.Name).To(Equal("meta_name"))
		Expect(compiled.Desc).To(Equal("meta_desc"))
	})

	It("converts runes into walls and floors", func() {
		dungeon := placeholder.NewMap(2, 1)

		dungeon.Grid[0][0] = &placeholder.Cell{Rune: '#'}
		dungeon.Grid[0][1] = &placeholder.Cell{Rune: ' '}

		meta := placeholder.Meta{}

		entities := placeholder.NewEntityCollection()

		compiled := Compile(&meta, dungeon, &entities)

		Expect(compiled.Tiles[0][0].Walkable).To(BeFalse())
		Expect(compiled.Tiles[0][1].Walkable).To(BeTrue())
	})

	It("adds entities to the map", func() {
		dungeon := placeholder.NewMap(0, 0)

		meta := placeholder.Meta{}

		entities := placeholder.NewEntityCollection()
		entities.AddMobs(placeholder.Mob{})
		entities.AddDoors(placeholder.Door{})
		entities.AddItems(placeholder.Item{})

		compiled := Compile(&meta, dungeon, &entities)

		Expect(len(compiled.Doors)).To(Equal(1))
		Expect(len(compiled.Mobs)).To(Equal(1))
		Expect(len(compiled.Items)).To(Equal(1))
	})

	It("adds mob loot", func() {
		dungeon := placeholder.NewMap(0, 0)
		meta := placeholder.Meta{}

		mob := placeholder.Mob{
			ParsedLoot: []placeholder.Loot{
				placeholder.Loot{
					Link:    "loot",
					MinHeld: 1,
					MaxHeld: 2,
				},
			},
		}

		loot := placeholder.Item{
			Link: "loot",
		}

		entities := placeholder.NewEntityCollection()
		entities.AddMobs(mob)
		entities.AddItems(loot)

		compiled := Compile(&meta, dungeon, &entities)

		Expect(len(compiled.Mobs)).To(Equal(1))
		Expect(len(compiled.Items)).To(Equal(1))

		Expect(len(compiled.Mobs[0].Loot)).To(Equal(1))
		Expect(compiled.Mobs[0].Loot[0].Link).To(Equal("loot"))
		Expect(compiled.Mobs[0].Loot[0].MinHeld).To(Equal(1))
		Expect(compiled.Mobs[0].Loot[0].MinHeld).To(Equal(2))

	})

	It("sets the spawn values for each cell", func() {
		dungeon := placeholder.NewMap(1, 1)
		dungeon.Grid[0][0] = &placeholder.Cell{
			Type: "mob",
			Link: "mob_link",
		}

		meta := placeholder.Meta{}

		entities := placeholder.NewEntityCollection()

		compiled := Compile(&meta, dungeon, &entities)

		Expect(compiled.Tiles[0][0].Spawn).To(Equal("mob_link"))
	})

	It("Adds the start position", func() {
		dungeon := placeholder.NewMap(0, 0)

		dungeon.StartPosition = &placeholder.Position{1, 1}

		meta := placeholder.Meta{}

		entities := placeholder.NewEntityCollection()

		compiled := Compile(&meta, dungeon, &entities)

		Expect(compiled.StartPosition).To(Equal(output.Position{1, 1}))
	})
})
