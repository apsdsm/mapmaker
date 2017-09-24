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

package placeholders

// Door is a placeholder for door entities in a level
type Door struct {

	// physical name
	Link string

	// true if door is locked
	Locked bool

	// link to key that unlocks/locks door
	Key string

	// events
	OnTry string
}

// NeedsKey returns true if the door requires a key
func (d *Door) NeedsKey() bool {
	return d.Key != ""
}
