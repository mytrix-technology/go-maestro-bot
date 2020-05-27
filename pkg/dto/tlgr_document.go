package dto

import "fmt"

type Document struct {
	FileID 		string	`json:"file_id"`
	FileName	string	`json:"file_name"`
}

func (d Document) String() string {
	return fmt.Sprintf("(file id: %s, file name: %s)", d.FileID, d.FileName)
}
