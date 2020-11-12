package console

import (
	"github.com/gogf/gf/frame/g"
	"github.com/spf13/cobra"
	"tools/app/models"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "迁移数据库文件",
	Run: func(cmd *cobra.Command, args []string) {
		err := models.DB.AutoMigrate(&models.Tools{})
		if err != nil {
			g.Log().Error(err)
		}
		g.Log().Println("迁移数据库成成功")
	},
}
