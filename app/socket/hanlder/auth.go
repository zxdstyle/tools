package hanlder

import "tools/app/socket"

func Login(client *socket.Client, message *socket.Message) {
	// @TODO 登录逻辑

	socket.ConnectionManager.SendMessage(client, message)
}
