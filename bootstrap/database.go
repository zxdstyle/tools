package bootstrap

import (
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tools/app/models"
)

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: g.Cfg().Get("database.link").(string),
	}), &gorm.Config{})
	if err != nil {
		g.Log().Error("数据库连接错误：" + err.Error())
	}

	if g.Cfg().GetBool("database.debug") {
		db.Debug()
	}

	models.DB = db
}
