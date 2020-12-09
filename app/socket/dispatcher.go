package socket

import "sync"

type FailedMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Message struct {
	Type    string `json:"type" v:"type@required#消息类型不能为空"`
	From    string `json:"name" v:""`
	Message string `json:"message"`
}

type HandlerFunc func(client *Client, message *Message)

var (
	handlers        = make(map[string]HandlerFunc)
	handlersRWMutex sync.RWMutex
)

// 注册websocket消息处理器
func Register(key string, handler HandlerFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()

	handlers[key] = handler

	return
}

// 获取对应的处理器
func getHandler(key string) (handler HandlerFunc, ok bool) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()

	handler, ok = handlers[key]

	return
}
