package file

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/apsdsm/mapmaker/formats/placeholder"
	yaml "gopkg.in/yaml.v2"
)

// EntityImporter is responsible for reading entities from an entity file and preparing them for import.
type EntityImporter struct {
	Entities *placeholder.EntityList
	Errors   *ErrorList

	// regular expressions
	matchLootItemMinMax  *regexp.Regexp // match `item[min..max]` notation
	matchLootItemZeroMax *regexp.Regexp // match `item[..max]` notation
	matchLootItemAmount  *regexp.Regexp // match `item[amount]` notation
	matchLootItemOne     *regexp.Regexp // match `item` notation
}

// NewEntityImporter initializes and returns a new EntityImporter.
func NewEntityImporter() *EntityImporter {
	i := EntityImporter{
		Entities: placeholder.NewEntityList(),
		Errors:   NewErrorList(),

		// e.g., copper_coin[2..10] <- give somewhere from 2 to 10 copper coins
		// matches: item_name, min, max
		matchLootItemMinMax: regexp.MustCompile(`^(.*)\[(\d+)\.\.(\d+)\]$`),

		// e.g., copper_coin[..10] <- give somewhere from 0 to 10 copper coins
		// matches: item_name, max
		matchLootItemZeroMax: regexp.MustCompile(`^(.*)\[\.\.(\d+)\]$`),

		// e.g., copper_coin[10] <- give exactly 10 copper coins
		// matches: item_name, amount
		matchLootItemAmount: regexp.MustCompile(`^(.*)\[(\d+)\]$`),

		// e.g., cooper_coin <- give one copper_coin
		// matches: item_name
		matchLootItemOne: regexp.MustCompile(`^(\w*)$`),
	}

	return &i
}

func (e *EntityImporter) Read(in string) error {
	absPath := absPath(in)
	bytes := readBytes(absPath)

	err := yaml.Unmarshal(bytes, &e.Entities)

	if err != nil {
		// add an error here
		return err
	}

	// unmarshaling the mob gave us most of its data for free, but we still need to parse a few strings for each one.
	for key, mob := range e.Entities.Mobs {

		// copy reference into mob for easy access
		mob.Reference = key

		// parse stuff that needs parsing
		e.parseLoot(mob)
	}

	return nil

}

// parse a loot string. Loot strings contain an item id, and a numerical range that describes the min/max carry of that
// item. There are a few was to write a loot string, so this method checks the string to discover the notation  before
// it starts buliding a loot placeholder.
func (e *EntityImporter) parseLoot(m *placeholder.Mob) {
	var matches []string
	var err error

	for _, raw := range m.Loot {

		var (
			item string
			min  string
			max  string
		)

		switch {

		// `loot[x..y]`
		case matchesRegExp(raw, e.matchLootItemMinMax, &matches):
			item = matches[1]
			min = matches[2]
			max = matches[3]

		// `loot[..x]`
		case matchesRegExp(raw, e.matchLootItemZeroMax, &matches):
			item = matches[1]
			min = "0"
			max = matches[2]

		// `loot[x]`
		case matchesRegExp(raw, e.matchLootItemAmount, &matches):
			item = matches[1]
			min = matches[2]
			max = matches[2]

		// `loot`
		case matchesRegExp(raw, e.matchLootItemOne, &matches):
			item = matches[1]
			min = "1"
			max = "1"
		}

		// min cannot be greater than max
		if min > max {
			e.Errors.Add(Error{
				Message:    fmt.Sprintf("mob: %s, loot string: '%s' => min cannot be greater than max", m.Reference, raw),
				FileName:   "",
				LineNumber: 0,
				IsWarning:  false,
			})
			return
		}

		loot := placeholder.Loot{}
		loot.Link = item

		// parse min value
		loot.MinHeld, err = strconv.Atoi(min)

		if err != nil {
			e.Errors.Add(Error{
				Message:    fmt.Sprintf("mob: %s : loot string: '%s' => min value could not be parsed as integer", m.Reference, raw),
				FileName:   "",
				LineNumber: 0,
				IsWarning:  false,
			})
		}

		// parse max value
		loot.MaxHeld, err = strconv.Atoi(max)

		if err != nil {
			e.Errors.Add(Error{
				Message:    fmt.Sprintf("mob: %s : loot string: '%s' => max value could not be parsed as integer", m.Reference, raw),
				FileName:   "",
				LineNumber: 0,
				IsWarning:  false,
			})
		}

		m.ParsedLoot = append(m.ParsedLoot, loot)
	}
}

// matches a regular expression. If the match is true, copies the submatches into the provided result array.
func matchesRegExp(in string, match *regexp.Regexp, result *[]string) bool {
	*result = match.FindStringSubmatch(in)

	if len(*result) > 0 {
		return true
	}

	return false
}
