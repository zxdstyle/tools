package console

import (
	"github.com/gogf/gf/frame/g"
	"github.com/spf13/cobra"
	"tools/app/socket"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "启动 API 服务",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	go socket.ConnectionManager.Start()

	g.Server().Run()
}
