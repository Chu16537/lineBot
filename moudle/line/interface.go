package line

import (
	"linebot/dao"
	"linebot/moudle/mongoDB"
	"linebot/moudle/web"

	"github.com/gin-gonic/gin"
)

type IHttp interface {
	PostLineBot(func(c *gin.Context))
}

type IDB interface {
	Query(userID string, time int64) []dao.Msg
	Save(data dao.Msg)
}

func NewHppt(http *web.Handler) IHttp {
	return http
}

func NewDB(db *mongoDB.Handler) IDB {
	return db
}
