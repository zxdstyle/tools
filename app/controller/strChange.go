package controller

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"tools/app/models"
	"tools/app/service/strService"
	"tools/app/support/h"
)

// json 字符串美化
func FormatJson(r *ghttp.Request) {
	str := r.Get("json")

	json, err := gjson.DecodeToJson(str)
	if err != nil {
		h.Failed(r, "json格式错误【"+err.Error()+"】")
	}

	jsonStr, err := json.ToJsonIndentString()
	if err != nil {
		h.Failed(r, "json格式错误")
	}

	r.Response.WriteJsonExit(jsonStr)
}

// json 字符串转结构体
func JsonToStruct(r *ghttp.Request) {

	structStr, err := strService.JsonToStruct(r.GetString("jsonStr"),
		r.GetString("modelName", "Test"),
		r.GetString("pkgName", "test"))
	if err != nil {
		h.Failed(r, err.Error())
	}

	r.Response.Write(structStr)
}

// Sql 字符串转结构体
func SqlToStruct(r *ghttp.Request) {

	modelStr, err := strService.DoGenModel(r.GetString("sqlStr"))
	if err != nil {
		h.Failed(r, err.Error())
	}

	r.Response.Write(modelStr)
}

func Language(r *ghttp.Request) {

	data := r.GetMap()

	params := make([]*models.Language, 0)
	for _, value := range data["data"].([]interface{}) {

		date := gtime.New(gconv.Int64(value.([]interface{})[0])).Format("Y-m-d H:i:s")

		ratings := value.([]interface{})[1].(float64)

		params = append(params, &models.Language{
			Name:    r.GetString("name"),
			Ratings: ratings,
			Date:    date,
		})

	}

	models.DB.Create(&params)

}

func GetAllLang(r *ghttp.Request) {
	var languages []*models.Language
	defaultLanguage := g.Array{"go", "php", "Java", "C", "C++", "Go", "Python"}
	if err := models.DB.Where("name IN ?", r.GetArray("languages", defaultLanguage)).
		Find(&languages).Error; err != nil {
		h.Failed(r, err.Error())
	}

	result := make(map[string][]interface{}, 0)
	for _, value := range languages {

		result[value.Name] = append(result[value.Name], g.Array{
			gtime.NewFromStr(value.Date).Format("Y-m-d"),
			value.Ratings})
	}

	res := make([]map[string]interface{}, 0)
	for name, data := range result {
		res = append(res, map[string]interface{}{
			"name": name,
			"data": data,
		})
	}

	h.Success(r, res)
}
