package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/frame/g"
	"tools/app/models"
)

var Enforcer *casbin.Enforcer

func init() {

	adapter, err := gormadapter.NewAdapterByDBUseTableName(models.DB, "", "rules")
	if err != nil {
		g.Log().Fatal(err)
	}

	enforcer, err := casbin.NewEnforcer("./config/casbin.conf", adapter)
	if err != nil {
		g.Log().Fatal(err)
	}

	Enforcer = enforcer
}
