package role

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"tools/app/models"
	"tools/app/support/h"
)

func filter(r *ghttp.Request) *gorm.DB {
	query := models.DB.Table("roles")

	name := r.GetString("name")
	if len(name) > 0 {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}

	slug := r.GetString("slug")
	if len(slug) > 0 {
		query.Where("slug LIKE ?", fmt.Sprintf("%%%s%%", slug))
	}

	status := r.GetInt64("status")
	if status > 0 {
		query.Where("status <> ?", status)
	}

	createdAt := r.GetArray("created_at")
	if len(createdAt) == 2 {
		start := carbon.Parse(createdAt[0]).StartOfDay()
		end := carbon.Parse(createdAt[1]).EndOfDay()

		query.Where("status BETWEEN ? AND ?", start, end)
	}

	return query
}

// 创建新角色
func CreateRole(request *CreateRoleRequest) error {
	err := models.DB.Create(&models.Roles{
		Name:   request.Name,
		Slug:   request.Slug,
		Status: request.Status,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

// 获取角色列表
func GetRoleList(r *ghttp.Request) *PaginatorResponse {
	pageSize := r.GetInt("pageSize", 20)
	page := r.GetInt("page", 1)

	var total int64
	filter(r).Count(&total)
	paginator := &PaginatorResponse{
		CurrentPage: page,
		PageSize:    pageSize,
		Total:       total,
	}

	offset := (paginator.CurrentPage - 1) * paginator.PageSize

	filter(r).Order("created_at DESC").Limit(paginator.PageSize).Offset(offset).Find(&paginator.Data)

	return paginator
}

// 编辑角色
func UpdateRole(r *ghttp.Request) {
	err := models.DB.Model(&models.Roles{}).Where("id", r.Get("role_id")).
		Select("name", "slug", "status").
		Updates(r.GetMap()).Error
	if err != nil {
		h.Failed(r, err.Error())
	}
}

// 删除角色
func DeleteRole(r *ghttp.Request) {
	err := models.DB.Where("id", r.Get("role_id")).Delete(&models.Roles{}).Error
	if err != nil {
		h.Failed(r, err.Error())
	}
}
