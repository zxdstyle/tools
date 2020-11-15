package strService

import (
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/gogf/gf/text/gstr"
	"strings"
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

// 生成结构体
func DoGenModel(schema string) (string, error) {
	stmt, err := sqlparser.Parse(schema)
	if err != nil {
		return "", err
	}

	var structStr strings.Builder
	statement := stmt.(*sqlparser.CreateTable)
	modelName := gstr.CamelCase(statement.NewName.Name.String())
	structStr.WriteString(fmt.Sprintf("type %s struct {\n", modelName))

	for _, column := range statement.Columns {
		structStr.WriteString(fmt.Sprintf("	%s	%s	`gorm:\"column:%s;%s\" json:\"%s\"` \n",
			gstr.CamelCase(column.Name),
			matchFieldType(column.Type),
			column.Name,
			matchOption(column.Options),
			column.Name))
	}
	structStr.WriteString("}")

	structStr.WriteString(fmt.Sprintf("\n\nfunc (*%s) TableName() string {\n	return \"%s\" \n}",
		modelName, statement.NewName.Name))

	return structStr.String(), nil
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

// 匹配字段配置
func matchOption(options []*sqlparser.ColumnOption) string {
	var resStr strings.Builder
	for _, option := range options {
		switch option.Type {
		case sqlparser.ColumnOptionNotNull:
			resStr.WriteString("NOT NULL;")
		case sqlparser.ColumnOptionComment:
			break
		case sqlparser.ColumnOptionDefaultValue:
			break
			//resStr.WriteString("default:"+fmt.Sprintf("%s", option.Value)+";")

		default:
			resStr.WriteString(option.Type.String() + ";")
		}

	}
	return resStr.String()
}
