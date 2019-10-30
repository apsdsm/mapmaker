package file

// Error holds data about errors encountered during the map import process
type Error struct {
	Message    string
	FileName   string
	LineNumber int
	IsWarning  bool
}

type ErrorList struct {
	errors []Error
}

func NewErrorList() *ErrorList {
	return &ErrorList{errors: make([]Error, 0)}
}

func (e *ErrorList) Add(error Error) {
	e.errors = append(e.errors, error)
}
