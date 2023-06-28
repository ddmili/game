package main

import (
	"game/api"
	"game/config"
	"game/internal/logger"
	"game/internal/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.InitConfig()

	logger.InitLogger(cfg.Debug)

	r := newFunction()

	r.Use(AccessJsMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.InitRouter(r.Group("/api"))

	go task.InitTask()

	r.Run(cfg.Port)

}

func newFunction() *gin.Engine {
	r := gin.Default()
	return r
}

// AccessJsMiddleware returns a middleware
func AccessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		//r:=c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Expose-Headers", "Authorization")
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST,GET")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		//删除缓存
		c.Set("content-type", "application/json")
		w.Header().Add("Cache-control", "no-cache,no-store,must-revalidate")
		w.Header().Add("Pragma", "no-cache")
		w.Header().Add("Expires", "0")
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
