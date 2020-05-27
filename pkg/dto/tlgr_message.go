package dto

import "fmt"

type Message struct {
	Text 		string 		`json:"text"`
	Chat 		Chat		`json:"chat"`
	Audio 		Audio		`json:"audio"`
	Voice 		Voice		`json:"voice"`
	Document 	Document 	`json:"Document"`
}

func (m Message) String() string {
	return fmt.Sprintf("(text: %s, audio: %s)", m.Text, m.Audio)
}