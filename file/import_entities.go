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

package file

import (
	"regexp"
	"strings"

	"path"

	"github.com/apsdsm/mapmaker/formats/placeholder"
)

// ImportEntities will import all entities in a make list and return a collection with those entities.
// Returns any errors or warnings that were encountered during the import process.
func ImportEntities(filePath string, errors []Error, warnings []Error) (*placeholder.EntityCollection, []Error, []Error) {
	absPath := absPath(filePath)
	bytes := readBytes(absPath)
	collection := placeholder.NewEntityCollection()
	makeList := placeholder.MakeList{}

	unmarshalYaml(*bytes, &makeList)

	for _, include := range makeList.Include {
		includePath := path.Join(path.Dir(absPath), include)
		errors, warnings = AddEntityToCollection(includePath, &collection, errors, warnings)
	}

	return &collection, errors, warnings
}

// AddEntityToCollection will add the entity stored in the file to the collection object
func AddEntityToCollection(path string, collection *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {
	absPath := absPath(path)
	bytes := readBytes(absPath)

	// we use the file extension to decide what kind of file it is
	if strings.Contains(path, ".mob.") {
		errors, warnings = addMobToCollection(bytes, path, collection, errors, warnings)
	} else if strings.Contains(path, ".door.") {
		addDoorToCollection(bytes, collection)
	} else if strings.Contains(path, ".key.") {
		addKeyToCollection(bytes, collection)
	} else if strings.Contains(path, ".item.") {
		errors, warnings = addItemToCollection(bytes, collection, errors, warnings)
	}
	return errors, warnings
}

// addMobToCollection adds a mob to the specified collection
func addMobToCollection(bytes *[]byte, resourcePath string, collection *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {
	mob := placeholder.Mob{}
	unmarshalYaml(*bytes, &mob)
	collection.AddMobs(mob)

	mob.ParsedLoot = make([]placeholder.Loot, len(mob.Loot))

	for _, raw := range mob.Loot {
		loot, err := parseLootString(raw)

		if err != nil {
			err.FileName = resourcePath
			errors = append(errors, *err)
		} else {
			mob.ParsedLoot = append(mob.ParsedLoot, loot)
		}
	}

	return errors, warnings
}

// parseLootString returns a loot object so long as the raw loot string is in a valid format
func parseLootString(raw string) (placeholder.Loot, *Error) {

	// the following regexps define a whitelist of valid item notations. Everything else is invalid.
	r1 := regexp.MustCompile(`^.*\[\d+~\d+\]$`) // item[min~max]
	r2 := regexp.MustCompile(`^.*\[\d+\]$`)     // item[amount]
	r3 := regexp.MustCompile(`^\w*$`)           // item, item123, item_foo

	if r1.MatchString(raw) {
		return placeholder.Loot{}, nil
	}

	if r2.MatchString(raw) {
		return placeholder.Loot{}, nil
	}

	if r3.MatchString(raw) {
		return placeholder.Loot{}, nil
	}

	return placeholder.Loot{}, &Error{
		LineNumber: -1,
		Message:    ("invalid item syntax: '" + raw + "'. Use notation: 'object_link[count]', 'object_link[min:max]', or 'object_link'"),
	}
}

// addDoorToCollection adds a door to the specified collection
func addDoorToCollection(bytes *[]byte, collection *placeholder.EntityCollection) {
	door := placeholder.Door{}
	unmarshalYaml(*bytes, &door)
	collection.Doors = append(collection.Doors, door)
}

// addKeyToCollection adds a key to the specified collection
func addKeyToCollection(bytes *[]byte, collection *placeholder.EntityCollection) {
	key := placeholder.Key{}
	unmarshalYaml(*bytes, &key)
	collection.Keys = append(collection.Keys, key)
}

// addItemToCollection adds an item to the specified collection
func addItemToCollection(bytes *[]byte, collection *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {
	item := placeholder.Item{}
	unmarshalYaml(*bytes, &item)
	collection.Items = append(collection.Items, item)

	return errors, warnings
}
