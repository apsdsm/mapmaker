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

package main

import (
	"fmt"
	"os"

	"github.com/apsdsm/mapmaker/app"
	"github.com/apsdsm/mapmaker/input"

	"github.com/gdamore/tcell"
)

func main() {
	screen := makeScreen()
	inputProvider := input.NewTcellProvider(screen)
	app := app.NewApp(screen, inputProvider)

	// ~~ should be application logic ~~
	screen.Clear()
	screen.Show()

	for app.Done != true {
		app.Update()
	}

	screen.Fini()
	os.Exit(0)
}

func makeScreen() tcell.Screen {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Println("could not create screen")
		os.Exit(1)
	}

	err = screen.Init()

	if err != nil {
		fmt.Println("could not init screen")
		os.Exit(1)
	}

	return screen
}
