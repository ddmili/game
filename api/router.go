package api

import (
	"game/api/websocket"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the route
func InitRouter(r *gin.RouterGroup) {
	r.GET("/websocket", websocket.HandleWebSocket)
}
