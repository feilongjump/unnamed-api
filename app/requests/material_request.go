package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type MaterialRequest struct {
	ManufacturerID uint64  `valid:"manufacturer_id" json:"manufacturer_id,omitempty"`
	No             string  `valid:"no" json:"no,omitempty"`
	Name           string  `valid:"name" json:"name,omitempty"`
	Spec           string  `valid:"-" json:"spec"`
	Category       string  `valid:"-" json:"category"`
	Unit           string  `valid:"-" json:"unit"`
	UnitPrice      float64 `valid:"-" json:"unit_price"`
	Remarks        string  `valid:"-" json:"remarks"`
}

func MaterialSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"manufacturer_id": []string{"required", "numeric", "existed:manufacturers"},
		"no":              []string{"required", "min_cn:2", "max_cn:30"},
		"name":            []string{"required", "min_cn:2", "max_cn:30"},
	}
	messages := govalidator.MapData{
		"manufacturer_id": []string{
			"required:厂家为必填项",
			"numeric:厂家必须是数字",
			"existed:此厂家不存在",
		},
		"no": []string{
			"required:编号为必填项",
			"min_cn:编号长度需至少 2 个字",
			"max_cn:编号长度不能超过 30 个字",
		},
		"name": []string{
			"required:姓名为必填项",
			"min_cn:姓名长度需至少 2 个字",
			"max_cn:姓名长度不能超过 30 个字",
		},
	}
	return validate(data, rules, messages)
}
