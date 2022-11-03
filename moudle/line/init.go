package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handler struct {
	bot  *linebot.Client
	http IHttp
	db   IDB
}

var handler *Handler

func Init(http IHttp, db IDB) {
	bot, err := linebot.New(CHANNEL_SECRET, CHANNEL_ACCESS_TOKEN)

	if err != nil {
		panic(err)
	}

	handler = &Handler{
		bot:  bot,
		http: http,
		db:   db,
	}

	handler.http.PostLineBot(handler.postLineBot)
}

func Get() *linebot.Client {
	return handler.bot
}
