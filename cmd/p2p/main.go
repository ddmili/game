package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-basic/uuid"
	"github.com/libp2p/go-reuseport"
)

// Client  is a client
type Client struct {
	UID     string
	Conn    net.Conn
	Address string
}

// Handler is a handler
type Handler struct {
	// 服务端句柄
	Listener net.Listener
	// 客户端句柄池
	ClientPool map[string]*Client
}

// Handle handles
func (s *Handler) Handle() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("获取连接句柄失败", err.Error())
			continue
		}
		id := uuid.New()
		s.ClientPool[id] = &Client{
			UID:     id,
			Conn:    conn,
			Address: conn.RemoteAddr().String(),
		}
		fmt.Println("一个客户端连接进去了,他的公网IP是", conn.RemoteAddr().String())
		// 暂时只接受两个客户端,多余的不处理
		if len(s.ClientPool) == 2 {
			// 交换双方的公网地址
			s.ExchangeAddress(id)
			break
		}
	}
}

// ExchangeAddress 交换地址
func (s *Handler) ExchangeAddress(id string) {
	for uid, client := range s.ClientPool {
		// 自己不交换
		if uid == id {
			continue
		}
		var data = make(map[string]string)
		data["dst_uid"] = client.UID     // 对方的 UID
		data["address"] = client.Address // 对方的公网地址
		body, _ := json.Marshal(data)
		if _, err := client.Conn.Write(body); err != nil {
			fmt.Println("交换地址时出现了错误", err.Error())
		}
	}
}

// Start start a program
func (s *Handler) Start() <-chan os.Signal {

	go s.Handle()
	fmt.Println("Program start...")
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("Program Exit...", s)
			return c
		case syscall.SIGUSR1:
			fmt.Println("usr1 signal", s)
			return c
		case syscall.SIGUSR2:
			fmt.Println("usr2 signal", s)
			return c
		}
	}

	fmt.Println("Program stop...")
	return c
}

func main() {
	listener, err := reuseport.Listen("tcp", "0.0.0.0:6999")
	if err != nil {
		panic("服务端监听失败" + err.Error())
	}
	h := &Handler{Listener: listener, ClientPool: make(map[string]*Client)}
	// 监听内网节点连接,交换彼此的公网 IP 和端口
	<-h.Start()

}
