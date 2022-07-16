package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ManufacturerRequest struct {
	No          string `valid:"no" json:"no,omitempty"`
	Name        string `valid:"name" json:"name,omitempty"`
	Category    string `valid:"category" json:"category"`
	PurchaserId uint64 `valid:"purchaser_id" json:"purchaser_id"`
	Address     string `valid:"-" json:"address"`
	Remarks     string `valid:"-" json:"remarks"`
}

func ManufacturerSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"no":           []string{"required", "min:2", "max:50"},
		"name":         []string{"required", "min_cn:2", "max_cn:30"},
		"category":     []string{"min_cn:2", "max_cn:50"},
		"purchaser_id": []string{"numeric", "existed:admin_users"},
	}
	messages := govalidator.MapData{
		"no": []string{
			"required:编号为必填项",
			"min_cn:编号长度需至少 2 个字",
			"max_cn:编号长度不能超过 50 个字",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
		},
		"category": []string{
			"min_cn:分类长度需至少 2 个字",
			"max_cn:分类长度不能超过 50 个字",
		},
		"purchaser_id": []string{
			"numeric:采购负责人必须是数字",
			"existed:此采购负责人不存在",
		},
	}
	return validate(data, rules, messages)
}
