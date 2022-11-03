package web

import (
	"linebot/conf"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
}

var handler *Handler

// 初始化

func Init() {
	handler = &Handler{
		engine: gin.Default(),
	}

	handler.pageInit()
}

func Get() *Handler {
	if handler.engine == nil {
		Init()
	}
	return handler
}

func Run() {
	if handler.engine == nil {
		Init()
	}
	handler.engine.Run(conf.Env.Port)
}

func (h *Handler) pageInit() {
	h.getPing()
}
