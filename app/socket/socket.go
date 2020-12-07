package socket

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WebSocketHandler struct {
	ws *ghttp.WebSocket
}

func Handle(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	go ConnectionManager.Start()
	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	client := NewClient(conn.RemoteAddr().String(), conn, uint64(time.Now().Unix()))

	go client.Read()
	go client.Write()

	// 触发用户连接事件
	ConnectionManager.Register <- client
}
