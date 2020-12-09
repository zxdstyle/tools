package socket

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"sync"
)

type Manager struct {
	Clients     map[string]*Client // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // appId+uuid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	//Login       chan *login        // 用户登录处理
	Unregister chan *Client // 断开连接处理程序
	Broadcast  chan []byte  // 广播 向全部成员发送数据
}

func NewConnectionManager() *Manager {
	return &Manager{
		Clients:  make(map[string]*Client),
		Users:    make(map[string]*Client),
		Register: make(chan *Client, 1000),
		//Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
}

func (m *Manager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.OnOpen(client)
		case client := <-m.Unregister:
			m.OnClose(client)
		}
	}
}

// 添加客户端
func (m *Manager) AddClients(client *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()

	m.Clients[client.UserId] = client
}

// 删除客户端
func (m *Manager) DelClients(client *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()

	if _, ok := m.Clients[client.UserId]; ok {
		delete(m.Clients, client.UserId)
	}
}

// 向客户端发送消息
func (m *Manager) SendMessage(client *Client, message interface{}) {
	msg, err := gjson.Encode(message)
	if err != nil {
		g.Log().Error(err)
	}

	client.SendData <- msg
}
