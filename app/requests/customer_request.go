package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CustomerRequest struct {
	Name           string `valid:"name" json:"name,omitempty"`
	Category       string `valid:"category" json:"category"`
	SalesmanId     uint64 `valid:"salesman_id" json:"salesman_id"`
	MerchandiserId uint64 `valid:"merchandiser_id" json:"merchandiser_id"`
	Grade          int8   `valid:"grade" json:"grade"`
	Currency       string `valid:"-" json:"currency"`
	PaymentMethod  string `valid:"-" json:"payment_method"`
	Address        string `valid:"-" json:"address"`
	Remarks        string `valid:"-" json:"remarks"`
}

func CustomerSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":            []string{"required", "min_cn:2", "max_cn:30"},
		"category":        []string{"min_cn:2", "max_cn:50"},
		"salesman_id":     []string{"numeric", "existed:admin_users"},
		"merchandiser_id": []string{"numeric", "existed:admin_users"},
		"grade":           []string{"numeric"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
		},
		"category": []string{
			"min_cn:分类长度需至少 2 个字",
			"max_cn:分类长度不能超过 50 个字",
		},
		"salesman_id": []string{
			"numeric:业务员必须是数字",
			"existed:此业务员不存在",
		},
		"merchandiser_id": []string{
			"numeric:业务员必须是数字",
			"existed:此跟单员不存在",
		},
		"grade": []string{
			"numeric:等级必须是数字",
		},
	}
	return validate(data, rules, messages)
}
