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
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

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

	err := yaml.Unmarshal(*bytes, &makeList)

	if err != nil {
		panic("error unmarshalling yaml" + err.Error())
	}

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
		errors, warnings = addDoorToCollection(bytes, path, collection, errors, warnings)
	} else if strings.Contains(path, ".item.") {
		errors, warnings = addItemToCollection(bytes, path, collection, errors, warnings)
	}
	return errors, warnings
}

// addMobToCollection adds a mob to the specified collection
func addMobToCollection(bytes *[]byte, resourcePath string, collection *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {
	mob := placeholder.Mob{}
	err := yaml.Unmarshal(*bytes, &mob)

	if err != nil {
		e := Error{
			LineNumber: -1,
			Message:    "unable to unmarshal yaml. Received error: " + err.Error(),
			FileName:   resourcePath,
		}

		errors := append(errors, e)

		return errors, warnings
	}

	mob.ParsedLoot = make([]placeholder.Loot, 0, len(mob.Loot))

	for _, raw := range mob.Loot {
		loot, err := parseLootString(raw, resourcePath)

		if err != nil {
			errors = append(errors, *err)
		} else {
			mob.ParsedLoot = append(mob.ParsedLoot, loot)
		}
	}

	collection.AddMobs(mob)

	return errors, warnings
}

// parseLootString returns a loot object so long as the raw loot string is in a valid format
func parseLootString(raw, resourcePath string) (placeholder.Loot, *Error) {

	// the following regexps define a whitelist of valid item notations. Everything else is invalid.
	r1 := regexp.MustCompile(`^(.*)\[(\d+)~(\d+)\]$`) // item[min~max]
	r2 := regexp.MustCompile(`^(.*)\[(\d+)\]$`)       // item[amount]
	r3 := regexp.MustCompile(`^(\w*)$`)               // item, item123, item_foo

	if r1.MatchString(raw) {
		matches := r1.FindStringSubmatch(raw)
		link := matches[1]
		min, _ := strconv.Atoi(matches[2])
		max, _ := strconv.Atoi(matches[3])
		return placeholder.NewLoot(raw, link, min, max), nil
	}

	if r2.MatchString(raw) {
		matches := r2.FindStringSubmatch(raw)
		link := matches[1]
		amount, _ := strconv.Atoi(matches[2])
		return placeholder.NewLoot(raw, link, amount, amount), nil
	}

	if r3.MatchString(raw) {
		matches := r3.FindStringSubmatch(raw)
		link := matches[1]
		return placeholder.NewLoot(raw, link, 1, 1), nil
	}

	return placeholder.Loot{}, &Error{
		LineNumber: -1,
		FileName:   resourcePath,
		Message:    "invalid item syntax: '" + raw + "'. Use notation: 'object_link[count]', 'object_link[min:max]', or 'object_link'",
	}
}

// addDoorToCollection adds a door to the specified collection
func addDoorToCollection(bytes *[]byte, resourcePath string, collection *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {
	door := placeholder.Door{}
	err := yaml.Unmarshal(*bytes, &door)

	if err != nil {
		e := Error{
			LineNumber: -1,
			Message:    "unable to unmarshal yaml. Received error: " + err.Error(),
			FileName:   resourcePath,
		}

		errors := append(errors, e)

		return errors, warnings
	}

	collection.Doors = append(collection.Doors, door)

	return errors, warnings
}

// addItemToCollection adds an item to the specified collection
func addItemToCollection(bytes *[]byte, resourcePath string, collection *placeholder.EntityCollection, errors []Error, warnings []Error) ([]Error, []Error) {
	item := placeholder.Item{}
	err := yaml.Unmarshal(*bytes, &item)

	if err != nil {
		e := Error{
			LineNumber: -1,
			Message:    "unable to unmarshal yaml. Received error: " + err.Error(),
			FileName:   resourcePath,
		}

		errors := append(errors, e)

		return errors, warnings
	}

	collection.Items = append(collection.Items, item)

	return errors, warnings
}
