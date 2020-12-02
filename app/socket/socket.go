package socket

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type WebSocketHandler struct {
	ws *ghttp.WebSocket
}

func (h *WebSocketHandler) Handle(r *ghttp.Request) {
	for {
		ws, err := r.WebSocket()
		if err != nil {
			g.Log().Error(err)
			return
		}

		h.ws = ws

		h.dispatch()
	}
}
