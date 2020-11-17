package role

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"tools/app/models"
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

	// 状态默认值
	if role.Status == 0 {
		role.Status = StatusActive
	}

	if err := gvalid.CheckStruct(role, nil); err != nil {
		h.Failed(r, err.FirstString(), 422)
	}
}

// 角色编辑
func ValidateUpdateRole(r *ghttp.Request) {
	rules := g.ArrayStr{
		"role_id@required|min:0",
		"status@in:1,2",
	}

	if err := gvalid.CheckMap(r.GetMap(), rules); err != nil {
		h.Failed(r, err.Error())
	}

	var count int64
	err := models.DB.Model(&models.Roles{}).Where("id", r.Get("role_id")).Count(&count).Error
	if err != nil {
		h.Failed(r, err.Error())
	}

	if count == 0 {
		h.Failed(r, "未找到该角色")
	}
}
