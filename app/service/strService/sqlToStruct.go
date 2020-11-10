package strService

import (
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/gogf/gf/text/gstr"
	"log"
)

var sqlTypeMap = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"bigint":             "int64",
	"int unsigned":       "uint",
	"integer unsigned":   "uint",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint32",
	"bigint unsigned":    "uint64",
	"bit":                "byte",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func DoGenModel(schema string) (structStr string) {
	stmt, err := sqlparser.Parse(schema)
	if err != nil  {
		log.Fatal(err)
	}

	structStr += fmt.Sprintf("type Test struct {\n")


	for _, column := range stmt.(*sqlparser.CreateTable).Columns {

		structStr += fmt.Sprintf("	%s	%s	`gorm:\"%s\" `json:\"%s\"` \n",
			gstr.CamelCase(column.Name),
			matchFieldType(column.Type),
			column.Name,
			column.Name)
	}

	structStr += "}"
fmt.Println(structStr)
	return
}

// 匹配字段类型
func matchFieldType(fieldType string) string {
	fieldType = gstr.ReplaceByMap(gstr.ToLower(fieldType), map[string]string{
		"(": "",
		")": "",
		"0": "",
		"1": "",
		"2": "",
		"3": "",
		"4": "",
		"5": "",
		"6": "",
		"7": "",
		"8": "",
		"9": "",
	})

	return sqlTypeMap[fieldType]
}