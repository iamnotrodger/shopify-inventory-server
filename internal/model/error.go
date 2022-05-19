package model

import "fmt"

type ErrInvalidModel struct {
	message string
}

func (e *ErrInvalidModel) Error() string {
	return fmt.Sprintf("invalid model: %v", e.message)
}
