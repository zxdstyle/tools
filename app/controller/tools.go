package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"tools/app/models"
	"tools/app/service/toolsService"
)

type ToolsController struct {
}

func (*ToolsController) Post(r *ghttp.Request) {
	var validator toolsService.CreateToolsValidator
	r.GetStruct(&validator)

	if err := validator.Validate(); err != nil {
		r.Response.WriteHeader(400)
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": err.(*gvalid.Error).FirstString(),
		})
	}

	tools := models.Tools{
		Name:        validator.Name,
		Icon:        validator.Icon,
		Description: validator.Description,
	}

	if err := models.DB.Create(&tools).Error; err != nil {
		r.Response.WriteHeader(400)
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": err.Error(),
		})
	}

	r.Response.WriteJson(g.Map{
		"code":    200,
		"message": "success",
	})
}
