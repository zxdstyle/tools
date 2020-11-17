package controller

import (
	"github.com/gogf/gf/net/ghttp"
	"tools/app/service/role"
	"tools/app/support/h"
)

// 创建新角色
func CreateRole(r *ghttp.Request) {
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
func GetRoleList(r *ghttp.Request) {
	h.Success(r, role.GetRoleList(r))
}

// 角色编辑
func UpdateRole(r *ghttp.Request) {
	role.ValidateUpdateRole(r)

	role.UpdateRole(r)

	h.Success(r)
}

// 删除角色
func DeleteRole(r *ghttp.Request) {
	role.DeleteRole(r)

	h.Success(r)
}
