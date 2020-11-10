package strService

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/text/gstr"
)

func JsonToStruct(jsonStr string) (structStr string, err error)  {
	json, err := gjson.DecodeToJson(jsonStr)
	if err != nil {
		return "", err
	}
	structStr += fmt.Sprintf("type Test struct {\n")

	structStr += parseField(json.ToMap())

	structStr += "}"
	return structStr, nil
}

func parseField(json map[string]interface{}) (structStr string)  {
	for key, value := range json {
		varType := fmt.Sprintf("%T", value)
		if varType == "<nil>" {
			varType = "interface{}"
		}
		structStr += fmt.Sprintf("	%s	%s	%s\n",
			gstr.CamelCase(key),
			varType,
			fmt.Sprintf("`json:\"%s\"`", key))
	}
	
	return structStr
}
