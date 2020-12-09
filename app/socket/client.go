package socket

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	SendData      chan []byte     // 待发送的数据
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接时间
	HeartbeatTime uint64          // 用户上次心跳时间
}

func NewClient(addr string, conn *websocket.Conn, firstTime uint64) *Client {
	return &Client{
		Addr:          addr,
		Socket:        conn,
		SendData:      make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
}

func (c *Client) Read() {
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			g.Log().Error(err)
		}

		fmt.Println("接收到数据：", string(msg))

		ConnectionManager.OnMessage(c, msg)
	}
}

func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		ConnectionManager.Unregister <- c
		c.Socket.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.SendData:
			if !ok {
				// 发送数据错误 关闭连接
				fmt.Println("Client发送数据 关闭连接", c.Addr, "ok", ok)
				return
			}

			fmt.Println("发送数据：", string(message))

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func handleMessage(c *Client, msg []byte) {
	var message Message
	if err := gjson.DecodeTo(msg, message); err != nil {
		error := &FailedMessage{
			Code:  BadRequest,
			Error: "消息格式错误",
		}
		ConnectionManager.SendMessage(c, error)
	}
}
