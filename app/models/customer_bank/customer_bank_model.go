package customer_bank

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/customer"
	"unnamed-api/pkg/database"
)

type CustomerBank struct {
	models.BaseModel

	CustomerID     uint64 `json:"customer_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Currency       string `json:"currency"`
	AccountName    string `json:"account_name"`
	AccountNumber  string `json:"account_number"`
	AccountBank    string `json:"account_bank"`
	BankAddress    string `json:"bank_address"`
	CompanyAddress string `json:"company_address"`

	// 通过 customer_id 关联用户
	Customer customer.Customer `json:"customer"`

	models.CommonTimestampsField
}

func (customerBank *CustomerBank) Create() {
	database.DB.Create(&customerBank)
}

func (customerBank *CustomerBank) Save() (rowsAffected int64) {
	result := database.DB.Save(&customerBank)
	return result.RowsAffected
}

func (customerBank *CustomerBank) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&customerBank)
	return result.RowsAffected
}
