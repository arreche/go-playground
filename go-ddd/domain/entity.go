package domain

import "fmt"

const (
	minNameLength = 3
	maxNameLength = 50
)

type Entity struct {
	ID   string
	Name string
}

func (e *Entity) Valid() error {
	if len(e.Name) < minNameLength || len(e.Name) > maxNameLength {
		return fmt.Errorf("name len must be between %d and %d characters", minNameLength, maxNameLength)
	}
	return nil
}
