package validator

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gvalid"
	"tools/app/models"
)

// `v:"unique:tableName,fieldName"`
func init() {

	const ruleName = "unique"

	gvalid.RegisterRule(ruleName, func(rule string, value interface{}, message string, params map[string]interface{}) error {
		ruleStr := gstr.Replace(rule, fmt.Sprintf("%s:", ruleName), "", 1)
		rules := gstr.Split(ruleStr, ",")
		if len(rules) < 2 {
			return errors.New("unique 规则缺失表名和字段名配置")
		}

		var count int64
		tableName := rules[0]
		fieldName := rules[1]

		if len(rules) == 2 {
			models.DB.Table(tableName).Where(fieldName, value).Count(&count)
			if count > 0 {
				return errors.New(fmt.Sprintf("%s 已被占用", fieldName))
			}
			return nil
		}

		return nil
	})
}
