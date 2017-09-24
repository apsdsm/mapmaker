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

		metaPlaceholders, mapPlaceholders := file.ImportMap(mapFile)

		fmt.Println("importing entities...")

		entityPlaceholders := file.ImportEntities(entitiesFile)

		fmt.Println("validating...")

		errors, warnings := file.ValidatePlaceholders(metaPlaceholders, mapPlaceholders, entityPlaceholders)

		fmt.Println("errors:")
		fmt.Println(len(errors))

		if len(errors) > 0 {
			for _, e := range errors {
				fmt.Println(e)
			}

			return cli.NewExitError("errors detected. Compilation did not finish.", 1)
		}

		fmt.Println("warnings:")
		fmt.Println(len(warnings))

		compiled := file.CompileLevel(metaPlaceholders, mapPlaceholders, entityPlaceholders)

		file.Write(compiled, output)

		fmt.Println("Done!")

		return nil
	}

	app.Run(os.Args)
}
