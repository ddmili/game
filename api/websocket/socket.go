package websocket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var myConn *websocket.Conn

// HandleWebSocket connects to the websocket server
func HandleWebSocket(c *gin.Context) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}
	defer conn.Close()

	myConn = conn
	// 循环读取客户端发送的消息
	for {
		_, message, err := myConn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		fmt.Printf("received: %s\n", message)

		// 发送消息给客户端
		err = myConn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("write error:", err)
			break
		}
	}
}

// WriteMessage writes a message
func WriteMessage(message string) error {
	if myConn == nil {
		return nil
	}
	err := myConn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println("write error:", err)
	}
	return err
}
