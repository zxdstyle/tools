package controller

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"tools/app/support/casbin"
)

type RoleController struct {
}

func (*RoleController) Post(r *ghttp.Request) {
	casbin.Enforcer.LoadPolicy()

	ok, _ := casbin.Enforcer.Enforce(r.GetString("user"), r.GetString("permission"), r.GetString("method"))
	fmt.Println(ok)

	casbin.Enforcer.AddPermissionForUser("user", "permission", "method")
}
