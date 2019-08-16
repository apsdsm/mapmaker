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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/apsdsm/mapmaker/file"
)

func TestMapImporter(t *testing.T) {

	meta, level := file.ImportDungeon("../fixtures/maps/meta_and_map.map")

	t.Run("it loads a metadata", func(t *testing.T) {
		assert.Equal(t, "prison", meta.Link)
		assert.Equal(t, meta.Name, "The Jovian Prison")
		assert.Equal(t, meta.Desc, "A gloomy building hidden deep in the Jovian woods.")
	})

	t.Run("it loads layout", func(t *testing.T) {
		assert.Equal(t, 11, level.Width)
		assert.Equal(t, 7, level.Height)
		assert.Equal(t, level.Width, len(level.Grid))

		for i := range level.Grid {
			assert.Equal(t, level.Height, len(level.Grid[i]))
		}
	})

	t.Run("it loads line metadata", func(t *testing.T) {
		assert.Equal(t, "mob", level.Grid[2][3].Type)
		assert.Equal(t, "mob1", level.Grid[2][3].Link)

		assert.Equal(t, "door", level.Grid[6][1].Type)
		assert.Equal(t, "door1", level.Grid[6][1].Link)

		assert.Equal(t, "waypoint", level.Grid[2][4].Type)
		assert.Equal(t, "waypoint1", level.Grid[2][4].Link)

		assert.Equal(t, "item", level.Grid[2][5].Type)
		assert.Equal(t, "item1", level.Grid[2][5].Link)
	})

	t.Run("it assigns a start position", func(t *testing.T) {
		assert.Equal(t, 2, level.StartPosition.X)
		assert.Equal(t, 1, level.StartPosition.Y)
	})

	t.Run("it sets rule value to ' ' if cell contains entity", func(t *testing.T) {
		assert.Equal(t, ' ', level.Grid[2][3].Rune, "mob entity rune should be blank")
		assert.Equal(t, ' ', level.Grid[6][1].Rune, "door entity rune should be blank")
		assert.Equal(t, ' ', level.Grid[2][4].Rune, "waypoint entity rune should be blank")
		assert.Equal(t, ' ', level.Grid[2][5].Rune, "item entity rune should be blank")
	})
}
