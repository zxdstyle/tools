package controller

import (
	"github.com/gogf/gf/net/ghttp"
	"tools/app/service/role"
	"tools/app/support/h"
)

type RoleController struct {
}

// 创建新角色
func (*RoleController) Post(r *ghttp.Request) {
	var request role.CreateRoleRequest
	if err := r.GetFormStruct(&request); err != nil {
		h.Failed(r, err.Error())
	}

	request.ValidateCreateRole(r)

	if err := role.CreateRole(&request); err != nil {
		h.Failed(r, err.Error(), 400)
	}

	h.Success(r)
}

// 获取角色列表
func (*RoleController) Get(r *ghttp.Request) {
	h.Success(r, role.GetRoleList(r))
}

func (*RoleController) Put(r *ghttp.Request) {

}
