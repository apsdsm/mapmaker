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

package placeholder

// Mob is a placeholder for mob entities in a level
type Mob struct {
	Name       string            `yaml:"name"`
	Reference  string            `yaml:"-"`
	Prototype  string            `yaml:"prototype"`
	Rune       string            `yaml:"rune"`
	Loot       []string          `yaml:"loot"`
	ParsedLoot []Loot            `yaml:"-"`
	Hp         int               `yaml:"hp"`
	Mp         int               `yaml:"mp"`
	Events     map[string]string `yaml:"events"`
}

// NeedsPrototype returns true if the mob requires a prototype entity
func (mob *Mob) NeedsPrototype() bool {
	return mob.Prototype != ""
}
