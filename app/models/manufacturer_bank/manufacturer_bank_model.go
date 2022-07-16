package manufacturer_bank

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/manufacturer"
	"unnamed-api/pkg/database"
)

type ManufacturerBank struct {
	models.BaseModel

	ManufacturerID uint64 `json:"manufacturer_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Currency       string `json:"currency"`
	AccountName    string `json:"account_name"`
	AccountNumber  string `json:"account_number"`
	AccountBank    string `json:"account_bank"`
	BankAddress    string `json:"bank_address"`
	CompanyAddress string `json:"company_address"`

	// 通过 manufacturer_id 关联用户
	Manufacturer manufacturer.Manufacturer `json:"manufacturer"`

	models.CommonTimestampsField
}

func (manufacturerBank *ManufacturerBank) Create() {
	database.DB.Create(&manufacturerBank)
}

func (manufacturerBank *ManufacturerBank) Save() (rowsAffected int64) {
	result := database.DB.Save(&manufacturerBank)
	return result.RowsAffected
}

func (manufacturerBank *ManufacturerBank) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&manufacturerBank)
	return result.RowsAffected
}
