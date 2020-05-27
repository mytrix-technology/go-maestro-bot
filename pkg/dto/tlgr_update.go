package dto

import "fmt"

type Update struct {
	UpdateID 	int 	`json:"updated_id"`
	Message 	Message	`json:"message"`
}

func (u Update) String() string  {
	return fmt.Sprintf("(update id: %d, message: %s)", u.UpdateID, u.Message)
}