package dto

import "fmt"

type Audio struct {
	FileID 		string 	`json:"file_id"`
	Duration 	int		`json:"duration"`
}

func (a Audio) String() string  {
	return fmt.Sprintf("(file id: %s, duration: %d)", a.FileID, a.Duration)
}
