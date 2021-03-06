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

package output

// Mob is a placeholder for mob entities in a level
type Mob struct {
	Name string
	Link string
	Prot string
	Rune string
	Loot []Loot
	Hp   string
	Mp   string
}

// NeedsPrototype returns true if the mob requires a prototype entity
func (mob *Mob) NeedsPrototype() bool {
	return mob.Prot != ""
}
