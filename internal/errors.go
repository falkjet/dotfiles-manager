package internal

import "fmt"

type TargetExists string

func (err TargetExists) Error() string {
	return fmt.Sprintf("Can't install at location %s. File already exists")
}

type MultipleFiles struct {
	files []string
}

func (err MultipleFiles) Error() string {
	return fmt.Sprintf("Cannot install file in same location as other files or directories: %v", err.files)
}
