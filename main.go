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

package main

import (
	"fmt"
	"os"

	"github.com/apsdsm/mapmaker/file"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Map Maker"
	app.Usage = "compile Valencia maps"
	app.Description = "Compile maps created in the Valencia format"

	var mapFile string
	var entitiesFile string
	var output string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "map, m",
			Usage:       "map file to import",
			Destination: &mapFile,
		},
		cli.StringFlag{
			Name:        "entities, e",
			Usage:       "makelist of entities to import",
			Destination: &entitiesFile,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "output location",
			Destination: &output,
		},
	}

	app.Action = func(c *cli.Context) error {

		warnings := make([]file.Error, 0, 10)

		// if no make file stop
		if entitiesFile == "" {
			return cli.NewExitError("no makefile supplied", 1)
		}

		// if no entities file stop
		if mapFile == "" {
			return cli.NewExitError("no entities makefile supplied", 1)
		}

		// if no output location stop
		if output == "" {
			return cli.NewExitError("no output location supplied", 1)
		}

		fmt.Println("importing map and meta...")

		errors := file.NewErrorList()

		dungeonImporter := file.NewDungeonImporter(file.DungeonImporterConfig{Errors: errors})

		err = dungeonImporter.Read(mapFile)

		if err != nil {
			return cli.NewExitError("error exit on import map", 1)
		}

		entityImporter := file.NewEntityImporter(file.EntityImporterConfig{Errors: errors})

		err = entityImporter.Read(entitiesFile)

		if err != nil {
			return cli.NewExitError("error exit on import entities", 1)
		}

		dungeonValidator := file.NewDungeonValidator(file.DungeonValidatorConfig{
			Errors: errors,
		})

		err = dungeonValidator.Validate(dungeonImporter.Dungeon, entityImporter.Entities)

		if err != nil {
			return cli.NewExitError("error exit on validate dungeon", 1)
		}

		compiler := file.NewCompiler(file.CompilerConfig{})

		// if this where it falls over?

		compiler.Compile(dungeonImporter.Dungeon, entityImporter.Entities)

		//metaPlaceholders, mapPlaceholders := file.ImportDungeon(mapFile)
		//
		//fmt.Println("importing entities...")
		//
		//entityPlaceholders, errors, warnings := file.ImportEntities(entitiesFile, errors, warnings)
		//
		//fmt.Println("validating...")
		//
		//errors, warnings = file.ValidateDungeon(metaPlaceholders, mapPlaceholders, entityPlaceholders)
		//
		//fmt.Println("errors:")
		//fmt.Println(len(errors))
		//
		//if len(errors) > 0 {
		//	for _, e := range errors {
		//		fmt.Println(e)
		//	}
		//
		//	return cli.NewExitError("errors detected. Compilation did not finish.", 1)
		//}
		//
		//fmt.Println("warnings:")
		//fmt.Println(len(warnings))
		//
		//compiled := file.Compile(metaPlaceholders, mapPlaceholders, entityPlaceholders)
		//
		//file.Write(compiled, output)

		fmt.Println("Done!")

		return nil
	}

	app.Run(os.Args)
}
