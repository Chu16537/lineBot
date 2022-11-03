package web

import (
	"github.com/gin-gonic/gin"
)

// lint bot
func (h *Handler) PostLineBot(fn func(c *gin.Context)) {
	h.engine.POST(PAGE_CALLBACK, fn)
}
