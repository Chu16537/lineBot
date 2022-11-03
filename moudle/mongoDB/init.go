package mongoDB

import (
	"linebot/conf"

	"gopkg.in/mgo.v2"
)

type Handler struct {
	session *mgo.Session
	db      *mgo.Database
}

var handler *Handler

// 初始化

func Init() {
	session, err := mgo.Dial(conf.Env.DBAddr)
	if err != nil {
		panic(err)
	}

	handler = &Handler{
		session: session,
		db:      session.DB("test"),
	}

}

func Get() *Handler {
	if handler.db == nil {
		Init()
	}

	return handler
}
