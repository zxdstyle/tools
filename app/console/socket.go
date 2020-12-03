package console

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/spf13/cobra"
	"net/http"
	"tools/app/socket"
	"tools/app/support/h"
)

var SocketCmd = &cobra.Command{
	Use:   "socket",
	Short: "启动 Socket 服务",
	Run: func(cmd *cobra.Command, args []string) {
		runSocket()
	},
}

func runSocket() {

	http.HandleFunc("/ws", socket.Handle)

	go socket.Start()

	port := g.Cfg().GetString("socket.Port")

	fmt.Println("WebSocket 启动程序成功", h.GetServerIp(), port)
	http.ListenAndServe(":"+port, nil)
}
