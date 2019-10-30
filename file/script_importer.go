package file

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type ScriptImporter struct {
	errs []string
}

func NewScriptImporter() *ScriptImporter {
	e := ScriptImporter{
		errs: make([]string, 0),
	}

	return &e
}

type state = int

const (
	searching = iota
	inBlockComment
	inLineComment
	reading
)

func (e *ScriptImporter) Read(in string) error {

	var err error

	// load file
	var path string
	var data *os.File

	if path, err = filepath.Abs(in); err != nil {
		return err
	}

	if data, err = os.Open(path); err != nil {
		return err
	}

	defer data.Close()

	// load data into a scanner
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	//inBlockComment = false
	readState := searching
	lineNumber := 1

	lineComment := regexp.MustCompile(`^\s*//`)           // -> //
	blockCommentStart := regexp.MustCompile(`^\s*/\*`)    // -> /*
	endBlockCommentStart := regexp.MustCompile(`^\s*\*/`) // -> */
	oneLineBlockComment := regexp.MustCompile(`^\s*/\*.*\*/`)

	// read file
	for scanner.Scan() {
		line := scanner.Text()

		// check for comments - should erase traces of the comment from the line
		switch {
		case lineComment.MatchString(line):
			fmt.Println("               comment: " + line)
			lineNumber++
			continue

		case oneLineBlockComment.MatchString(line):
			fmt.Println("one line block comment: " + line)
			lineNumber++
			continue

		case blockCommentStart.MatchString(line):
			fmt.Println("  opened block comment: " + line)
			readState = inBlockComment
		}

		// after comments have been handled, we parse what's left of the line
		switch readState {
		case searching:
			fmt.Println("    interpretable line: " + line)

		case inBlockComment:
			if endBlockCommentStart.MatchString(line) {
				fmt.Println("  closed block comment: " + line)
				readState = searching
			} else {
				fmt.Println("  inside block comment: " + line)
			}
		}

		lineNumber++
	}

	return nil
}
