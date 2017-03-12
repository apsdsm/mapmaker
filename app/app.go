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

package app

import "github.com/apsdsm/mapmaker/input"

// An App is an instance of the Map Maker applicaton
type App struct {
	screen Screen
	input  input.Provider
	Done   bool
}

// NewApp returns a new application instance
func NewApp(screen Screen, input input.Provider) *App {
	a := App{}
	a.screen = screen
	a.input = input
	a.Done = false

	return &a
}

// HandleInput will provide the app with a key press.
func (a *App) HandleInput() {
	key := a.input.GetInput()

	switch key {
	case input.Quit:
		a.Done = true
	}
}

// Update runs the app through a single update cycle
func (a *App) Update() {
	a.HandleInput()
}
