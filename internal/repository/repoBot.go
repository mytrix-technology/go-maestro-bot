package repository

import (
	"encoding/json"
	"github.com/mytrix-technology/go-maestro-bot/pkg/dto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

//Define a few constants and variable to handle different commands
const startCommand string = "/start"
var lenStartCommand int = len(startCommand)

const punchCommand string = "/punch"
var lenPunchCommand int = len(punchCommand)

const botTag string = "@yudz_maestro_bot"
var lenBotTag int = len(botTag)

//Pass token and sensible APIs through environment variables
const telegramApiBaseUrl string = "https://api.telegram.org/bot"
const telegramApiSendMessage string = "/sendMessage"
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"
var telegramApi string = telegramApiBaseUrl + os.Getenv(telegramTokenEnv) + telegramApiSendMessage

const lyricsApiEnv string = "LYRICS_API"
var lyricsApi string = os.Getenv(lyricsApiEnv)

//Sanitize remove clutter like /start /punch or the bot name from the string s passed as input
func Sanitize(s string) string {
	if len(s) >= lenStartCommand {
		if s[:lenStartCommand] == startCommand {
			s = s[lenStartCommand:]
		}
	}

	if len(s) >= lenPunchCommand {
		if s[:lenPunchCommand] == punchCommand {
			s = s[lenPunchCommand:]
		}
	}

	if len(s) >= lenBotTag {
		if s[:lenBotTag] == botTag {
			s = s[lenBotTag:]
		}
	}

	return s
}

// getPunchline calls the RapLyrics API to get a punchline back.
func GetPunchline(seed string) (string, error) {
	lyricsResp, err := http.PostForm(
		lyricsApi,
		url.Values{"input": {seed}})
	if err != nil {
		log.Printf("error while calling lyrics %s", err.Error())
		return "", err
	}

	var punchLine dto.Lyric
	if err := json.NewDecoder(lyricsResp.Body).Decode(&punchLine); err != nil {
		log.Printf("could not decode incoming punchline %s", err.Error())
		return "", err
	}

	defer lyricsResp.Body.Close()
	return punchLine.Punch, nil
}

func SendTextTelegramChat(chatID int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatID)
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text": {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of telegram response: %s", bodyString)

	return bodyString, nil
}
