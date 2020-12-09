package socket

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
)

// 事件处理

// 监听WebSocket连接关闭事件
func (m *Manager) OnClose(client *Client) {
	m.DelClients(client)
}

// 监听WebSocket连接打开事件
func (m *Manager) OnOpen(client *Client) {
	m.AddClients(client)

	message := &Message{
		Type: "system",
		From: "system",
	}

	m.SendMessage(client, message)
}

// 监听WebSocket消息事件
func (m *Manager) OnMessage(client *Client, msg []byte) {
	message := &Message{}
	if err := gjson.DecodeTo(msg, message); err != nil {
		errorMsg := &FailedMessage{
			Code:  BadRequest,
			Error: fmt.Sprintf("消息格式错误:%s", err.Error()),
		}
		m.SendMessage(client, errorMsg)
		return
	}

	if handler, ok := getHandler(message.Type); ok {
		handler(client, message)
	} else {
		error := &FailedMessage{
			Code:  NotFound,
			Error: "路由不存在",
		}
		m.SendMessage(client, error)
		return
	}
}
