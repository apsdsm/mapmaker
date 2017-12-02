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
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/apsdsm/mapmaker/formats/placeholder"
)

// Read will read a file as a map
func Read(path string) *placeholder.Map {
	var m placeholder.Map
	file, _ := ioutil.ReadFile(path)
	json.Unmarshal(file, &m)
	return &m
}

// readBytes will read and return the bytes from a file path and panic if it can't
func readBytes(filePath string) *[]byte {
	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		panic("error while opening file: " + err.Error())
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		panic("error while reading data from file: " + err.Error())
	}

	return &bytes
}
