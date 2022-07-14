package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CustomerContactRequest struct {
	CustomerID uint64 `valid:"customer_id" json:"customer_id,omitempty"`
	Name       string `valid:"name" json:"name,omitempty"`
	Phone      string `valid:"-" json:"phone,omitempty"`
	Email      string `valid:"email" json:"email"`
	Fax        string `valid:"-" json:"fax"`
	IsDefault  uint8  `valid:"is_default" json:"is_default"`
}

func CustomerContactSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"customer_id": []string{"required", "numeric", "existed:customers"},
		"name":        []string{"required", "min_cn:2", "max_cn:30"},
		"email":       []string{"min:4", "max:30", "email"},
		"is_default":  []string{"numeric"},
	}
	messages := govalidator.MapData{
		"customer_id": []string{
			"required:客户为必填项",
			"numeric:客户必须是数字",
			"existed:此客户不存在",
		},
		"name": []string{
			"required:姓名为必填项",
			"min_cn:姓名长度需至少 2 个字",
			"max_cn:姓名长度不能超过 30 个字",
		},
		"email": []string{
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"is_default": []string{
			"numeric:是否为默认联系人必须是数字",
		},
	}
	return validate(data, rules, messages)
}
