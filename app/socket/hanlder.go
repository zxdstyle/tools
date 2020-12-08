package socket

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

var (
	ConnectionManager = NewConnectionManager()
)

func Handler(r *ghttp.Request) {
	connection, err := r.WebSocket()
	if err != nil {
		g.Log().Error(err)
	}

	address := connection.Conn.RemoteAddr().String()
	fmt.Println("webSocket 建立连接:", address)
	currentTime := uint64(time.Now().Unix())
	client := NewClient(address, connection.Conn, currentTime)

	go client.Read()
	go client.Write()

	// 触发用户连接事件
	ConnectionManager.Register <- client
}
