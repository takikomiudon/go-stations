package model

import "fmt"

type ErrNotFound struct {
	ID int
}


func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("TODO with ID %d not found", e.ID)
}
