package http

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
/*func HandleTelegramWebhook(w http.ResponseWriter, r *http.Request) {
	//Parse Incoming request
	var update, err = ParseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	//Sanitize Input
	var sanitizedSeed = repository.Sanitize(update.Message.Text)

	//Call Lyric to get a puncline
	var lyric, errLyrics = repository.GetPunchline(sanitizedSeed)
	if errLyrics != nil {
		log.Printf("got error when calling Lyrics API %s", errLyrics.Error())
		return
	}

	// Send the punchline back to Telegram
	var telegramResponseBody, errTelegram = repository.SendTextTelegramChat(update.Message.Chat.ID, lyric)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("punchline %s succesfully distributed to chat id %d", lyric, update.Message.Chat.ID)
	}
}*/