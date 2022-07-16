package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CustomerBankRequest struct {
	CustomerID     uint64 `valid:"customer_id" json:"customer_id,omitempty"`
	Name           string `valid:"name" json:"name,omitempty"`
	Currency       string `valid:"-" json:"currency"`
	AccountName    string `valid:"-" json:"account_name"`
	AccountNumber  string `valid:"-" json:"account_number"`
	AccountBank    string `valid:"-" json:"account_bank"`
	BankAddress    string `valid:"-" json:"bank_address"`
	CompanyAddress string `valid:"-" json:"company_address"`
}

func CustomerBankSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"customer_id": []string{"required", "numeric", "existed:customers"},
		"name":        []string{"required", "min_cn:2", "max_cn:30"},
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
	}
	return validate(data, rules, messages)
}
