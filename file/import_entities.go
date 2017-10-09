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
	"strings"

	"path"

	"github.com/apsdsm/mapmaker/placeholders"
)

// ImportEntities will import all entities in a make list and return a collection with those entities
func ImportEntities(filePath string) *placeholders.EntityCollection {
	absPath := absPath(filePath)
	bytes := readBytes(absPath)
	collection := placeholders.NewEntityCollection()
	makeList := placeholders.MakeList{}

	unmarshalYaml(*bytes, &makeList)

	for _, include := range makeList.Include {
		includePath := path.Join(path.Dir(absPath), include)
		AddEntityToCollection(includePath, collection)
	}

	return collection
}

// AddEntityToCollection will add the entity stored in the file to the collection object
func AddEntityToCollection(path string, collection *placeholders.EntityCollection) {
	absPath := absPath(path)
	bytes := readBytes(absPath)

	// we use the file extension to decide what kind of file it is
	if strings.Contains(path, ".mob.") {
		addMobToCollection(bytes, collection)

	} else if strings.Contains(path, ".door.") {
		addDoorToCollection(bytes, collection)

	} else if strings.Contains(path, ".key.") {
		addKeyToCollection(bytes, collection)
	}
}

// addMobToCollection adds a mob to the specified collection
func addMobToCollection(bytes *[]byte, collection *placeholders.EntityCollection) {
	mob := placeholders.Mob{}
	unmarshalYaml(*bytes, &mob)
	collection.Mobs = append(collection.Mobs, mob)
}

// addDoorToCollection adds a door to the specified collection
func addDoorToCollection(bytes *[]byte, collection *placeholders.EntityCollection) {
	door := placeholders.Door{}
	unmarshalYaml(*bytes, &door)
	collection.Doors = append(collection.Doors, door)
}

// addKeyToCollection adds a key to the specified collection
func addKeyToCollection(bytes *[]byte, collection *placeholders.EntityCollection) {
	key := placeholders.Key{}
	unmarshalYaml(*bytes, &key)
	collection.Keys = append(collection.Keys, key)
}
