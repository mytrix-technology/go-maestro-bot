package http

import (
	"encoding/json"
	"errors"
	"github.com/mytrix-technology/go-maestro-bot/pkg/dto"
	"log"
	"net/http"
)

func ParseTelegramRequest(r *http.Request) (*dto.Update, error)  {
	var update dto.Update
	
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}

	if update.UpdateID == 0 {
		log.Printf("invalid update id, got update id = 0")
		return nil, errors.New("invalid update id of 0 indicates failure to paste incoming update")
	}

	return &update, nil
}