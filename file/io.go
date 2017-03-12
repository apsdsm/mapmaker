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

package file

import (
	"encoding/json"
	"io/ioutil"

	"github.com/apsdsm/mapmaker/maps"
)

// Out will output a map as a file
func Out(m *maps.Map, path string) {
	marshalled, _ := json.Marshal(m)

	ioutil.WriteFile(path, marshalled, 0664)
}

// In will read a file as a map
func In(path string) *maps.Map {
	var m maps.Map

	file, _ := ioutil.ReadFile(path)

	json.Unmarshal(file, &m)

	return &m
}
