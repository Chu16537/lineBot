package mongoDB

import (
	"gopkg.in/mgo.v2"
)

type Handler struct {
	session *mgo.Session
	db      *mgo.Database
}

var handler *Handler

// Connection URI
const uri = "127.0.0.1:27017"

// 初始化

func Init() {
	session, err := mgo.Dial(uri)
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
