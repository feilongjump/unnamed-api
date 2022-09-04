package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"unnamed-api/app/models"
)

type ExhibitRequest struct {
	ManufacturerID uint64      `valid:"manufacturer_id" json:"manufacturer_id"`
	Mo             string      `valid:"mo" json:"mo"`
	No             string      `valid:"no" json:"no"`
	Name           models.Lang `valid:"name" json:"name"`
	Spec           string      `valid:"-" json:"spec"`
	Series         string      `valid:"-" json:"series"`
	Material       string      `valid:"-" json:"material"`
	UnitPrice      float64     `valid:"-" json:"unit_price"`
	QuotedPrice    float64     `valid:"-" json:"quoted_price"`
	TaxRebateRate  float64     `valid:"-" json:"tax_rebate_rate"`
	Describe       models.Lang `valid:"-" json:"describe"`
	PackDescribe   models.Lang `valid:"-" json:"pack_describe"`
	Remarks        string      `valid:"-" json:"remarks"`
}

func ExhibitSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"manufacturer_id": []string{"required", "numeric", "existed:manufacturers"},
		"mo":              []string{"required", "min_cn:2", "max_cn:30"},
		"no":              []string{"required", "min_cn:2", "max_cn:30"},
		"name":            []string{"required"},
	}
	messages := govalidator.MapData{
		"manufacturer_id": []string{
			"required:厂家为必填项",
			"numeric:厂家必须是数字",
			"existed:此厂家不存在",
		},
		"mo": []string{
			"required:MO 为必填项",
			"min_cn:MO 长度需至少 2 个字",
			"max_cn:MO 长度不能超过 30 个字",
		},
		"no": []string{
			"required:编号为必填项",
			"min_cn:编号长度需至少 2 个字",
			"max_cn:编号长度不能超过 30 个字",
		},
		"name": []string{
			"required:名称为必填项",
			"json:名称应为 json 类型字段",
		},
	}
	return validate(data, rules, messages)
}
