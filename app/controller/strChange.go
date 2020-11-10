package controller

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"log"
	"tools/app/service/strService"
)

// json 字符串美化
func FormatJson(r *ghttp.Request) {
	str := r.Get("json")

	json, err := gjson.DecodeToJson(str)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code": 400,
			"message": "json格式错误【"+err.Error()+"】",
		})
	}

	jsonStr, err := json.ToJsonIndentString()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code": 400,
			"message": "json格式错误",
		})
	}

	r.Response.WriteJsonExit(jsonStr)
}

// json 字符串转结构体
func JsonToStruct(r *ghttp.Request)  {
	str := r.Get("json")

	structStr, err := strService.JsonToStruct(str.(string))
	if err != nil {
		r.Response.Header().Set("status", "400")
		r.Response.WriteJsonExit(g.Map{
			"code": 400,
			"message": err.Error(),
		})
	}

	r.Response.Write(structStr)
}

func SqlToStruct(r *ghttp.Request) {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}


	strService.DoGenModel(string(bytes[:]))
}