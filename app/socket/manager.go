package socket

import (
	"fmt"
	"sync"
)

type Manager struct {
	Clients     map[*Client]bool   // 全部的连接
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
		Clients:  make(map[*Client]bool),
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
			m.EventRegister(client)
			fmt.Println("触发链接事件")
		case client := <-m.Unregister:
			m.EventUnRegister(client)
		}
	}
}

func (m *Manager) EventUnRegister(client *Client) {
	m.DelClients(client)
}

func (m *Manager) EventRegister(client *Client) {
	m.AddClients(client)

	client.SendData <- []byte("连接成功")
}

// 添加客户端
func (m *Manager) AddClients(client *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()

	m.Clients[client] = true
}

// 删除客户端
func (m *Manager) DelClients(client *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()

	if _, ok := m.Clients[client]; ok {
		delete(m.Clients, client)
	}
}
