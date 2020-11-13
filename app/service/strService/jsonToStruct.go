package strService

import (
	"errors"
	"github.com/ChimeraCoder/gojson"
	"strings"
)

// JSON 数据生成结构体
func JsonToStruct(jsonStr string, structName string, pkgName string) (string, error) {
	// 结构需要生成的 tag
	tags := make([]string, 0)
	tags = append(tags, "json")

	structBytes, err := gojson.Generate(strings.NewReader(jsonStr),
		gojson.ParseJson,
		structName,
		pkgName,
		tags, true, true)
	if err != nil {
		return "", errors.New("JSON 格式错误")
	}

	return string(structBytes), nil
}
