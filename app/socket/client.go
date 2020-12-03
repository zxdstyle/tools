package socket

import "github.com/gorilla/websocket"

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

func (*Client) Read() {

}

func (*Client) Write() {

}
