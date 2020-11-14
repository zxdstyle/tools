package controller

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"log"
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
	fmt.Println(r.GetString("modelName", "Test"))
	r.Response.Write(structStr)
}

func SqlToStruct(r *ghttp.Request) {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	strService.DoGenModel(string(bytes[:]))
}
