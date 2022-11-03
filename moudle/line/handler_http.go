package line

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (h *Handler) postLineBot(c *gin.Context) {

	events, err := h.bot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}
	for _, event := range events {

		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				userId := event.Source.UserID
				token := event.ReplyToken
				msg := linebot.NewTextMessage(message.Text)

				replyText := h.messageAnalyze(userId, token, msg.Text)

				if _, err = h.bot.ReplyMessage(token, linebot.NewTextMessage(replyText)).Do(); err != nil {
					log.Print(err)
				}
			}
		case linebot.EventTypeJoin:
			fmt.Printf("[postLineBot] EventTypeJoin %v", event.Members)

		default:
			return
		}

	}

}
