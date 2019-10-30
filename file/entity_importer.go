package file

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/apsdsm/mapmaker/formats/placeholder"
	yaml "gopkg.in/yaml.v2"
)

type EntityImporter struct {
	Entities *placeholder.EntityList
	Errors   *ErrorList

	// regular expressions
	matchLootItemMinMax  *regexp.Regexp // match `item[min..max]` notation
	matchLootItemZeroMax *regexp.Regexp // match `item[..max]` notation
	matchLootItemAmount  *regexp.Regexp // match `item[amount]` notation
	matchLootItemOne     *regexp.Regexp // match `item` notation
}

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
		return err
	}

	// parse mobs
	for key, mob := range e.Entities.Mobs {

		// copy reference into mob for easy access
		mob.Reference = key

		if err = e.parseLootStrings(mob); err != nil {
			return err
		}

	}

	return nil

}

func (e *EntityImporter) parseLootStrings(m *placeholder.Mob) (err error) {

	// copy the parse loot logic in here
	var matches []string

	for _, raw := range m.Loot {
		switch {

		// `loot[x..y]`
		case matchesRegExp(raw, e.matchLootItemMinMax, &matches):
			if err = e.parseLootVars(m, matches[1], matches[2], matches[3], raw); err != nil {
				return
			}

		// `loot[..x]`
		case matchesRegExp(raw, e.matchLootItemZeroMax, &matches):
			if err = e.parseLootVars(m, matches[1], "0", matches[2], raw); err != nil {
				return
			}

		// `loot[x]`
		case matchesRegExp(raw, e.matchLootItemAmount, &matches):
			if err = e.parseLootVars(m, matches[1], matches[2], matches[2], raw); err != nil {
				return
			}

		// `loot`
		case matchesRegExp(raw, e.matchLootItemOne, &matches):
			if err = e.parseLootVars(m, matches[1], "1", "1", raw); err != nil {
				return
			}
		}
	}

	return nil
}

// matches a regular expression. If the match is true, copies the submatches into the provided result array.
func matchesRegExp(in string, match *regexp.Regexp, result *[]string) bool {
	*result = match.FindStringSubmatch(in)

	if len(*result) > 0 {
		return true
	}

	return false
}

func (e *EntityImporter) parseLootVars(m *placeholder.Mob, item, min, max, raw string) (err error) {

	// min cannot be greater than max
	if min > max {
		e.Errors.Add(Error{
			Message:    fmt.Sprintf("mob: %s, loot string: '%s' => min cannot be greater than max", m.Reference, raw),
			FileName:   "",
			LineNumber: 0,
			IsWarning:  false,
		})
		return nil
	}

	loot := placeholder.Loot{}
	loot.Link = item
	loot.MinHeld, err = strconv.Atoi(min)

	if err != nil {
		e.Errors.Add(Error{
			Message:    fmt.Sprintf("mob: %s : loot string: '%s' => min value could not be parsed as integer", m.Reference, raw),
			FileName:   "",
			LineNumber: 0,
			IsWarning:  false,
		})
	}

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

	return nil
}
