package handler

import (
	http2 "github.com/mytrix-technology/go-maestro-bot/internal/http"
	"github.com/mytrix-technology/go-maestro-bot/internal/repository"
	"log"
	"net/http"
)

func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// Parse incoming request
	var update, err = http2.ParseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// Sanitize input
	var sanitizedSeed = repository.Sanitize(update.Message.Text)

	// Call RapLyrics to get a punchline
	var lyric, errRapLyrics = repository.GetPunchline(sanitizedSeed)
	if errRapLyrics != nil {
		log.Printf("got error when calling RapLyrics API %s", errRapLyrics.Error())
		return
	}

	// Send the punchline back to Telegram
	var telegramResponseBody, errTelegram = repository.SendTextTelegramChat(update.Message.Chat.ID, lyric)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("punchline %s successfully distributed to chat id %d", lyric, update.Message.Chat.ID)
	}
}
