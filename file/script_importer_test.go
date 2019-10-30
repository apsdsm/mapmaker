package file_test

import (
	"testing"

	"github.com/apsdsm/mapmaker/file"
)

func TestEntityImporter_Import(t *testing.T) {
	t.Run("it imports a mob", func(t *testing.T) {

		importer := file.NewScriptImporter()

		err := importer.Read("../fixtures/entities/entity_test.vals")

		if err != nil {
			panic(err)
		}

	})
}
