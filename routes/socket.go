package routes

import (
	"tools/app/socket"
	"tools/app/socket/hanlder"
)

func initWebSocketRoute() {
	socket.Register("login", hanlder.Login)
}
