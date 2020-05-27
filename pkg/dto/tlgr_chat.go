package dto

import "fmt"

type Chat struct {
	ID int `json:"id"`
}

func (c Chat) String() string {
	return fmt.Sprintf("(id: %d)", c.ID)
}