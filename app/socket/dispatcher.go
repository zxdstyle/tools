package socket

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type FailedMessage struct {
	Type  string `json:"type" v:"type@required#消息类型不能为空"`
	Error string `json:"data" v:""`
}

type Message struct {
	Type string `json:"type" v:"type@required#消息类型不能为空"`
	From string `json:"name" v:""`
	To   string `json:"to"`
}

// 分发 websocket 请求
func (h *WebSocketHandler) dispatch() {
	_, msg, err := h.ws.ReadMessage()
	if err != nil {
		g.Log().Error(err)
	}

	message := h.parseMessage(msg)
}

func (h *WebSocketHandler) parseMessage(msg []byte) (message *Message) {
	if err := gjson.DecodeTo(msg, message); err != nil {
		h.writeFailed(err.Error())
		return nil
	}
	if e := gvalid.CheckStruct(message, nil); e != nil {
		h.writeFailed(e.FirstString())
		return
	}
	return message
}

func (h *WebSocketHandler) writeFailed(errMsg string) error {

	msg := FailedMessage{
		Type:  "system",
		Error: errMsg,
	}

	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}

	return h.ws.WriteMessage(ghttp.WS_MSG_TEXT, b)
}

func (h *WebSocketHandler) write(msg Message) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}

	return h.ws.WriteMessage(ghttp.WS_MSG_TEXT, b)
}
