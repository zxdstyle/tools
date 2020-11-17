package role

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"tools/app/support/h"
)

const (
	StatusActive  = 1 // 启用
	StatusInvalid = 2 // 禁用
)

type CreateRoleRequest struct {
	Name   string `v:"name@required|unique:roles,name"`
	Slug   string `v:"slug@required|unique:roles,slug"`
	Status int    `v:"in:1,2"`
}

// 创建角色验证器
func (role *CreateRoleRequest) ValidateCreateRole(r *ghttp.Request) {
	if err := r.GetStruct(&role); err != nil {
		h.Failed(r, err.Error(), 500)
	}

	// 状态默认值
	if role.Status == 0 {
		role.Status = StatusActive
	}

	if err := gvalid.CheckStruct(role, nil); err != nil {
		h.Failed(r, err.FirstString(), 422)
	}
}

type UpdateRoleRequest struct {
	Name   string `v:"name@unique:roles,name,"`
	Slug   string `v:"slug@unique:roles,slug"`
	Status int    `v:"status@in:1,2"`
}
