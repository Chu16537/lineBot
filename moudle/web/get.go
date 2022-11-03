package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getPing() {
	h.engine.GET(PAGE_PING, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
