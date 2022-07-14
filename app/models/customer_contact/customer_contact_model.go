package customer_contact

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/customer"
	"unnamed-api/pkg/database"
)

type CustomerContact struct {
	models.BaseModel

	CustomerID uint64 `json:"customer_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Fax        string `json:"fax"`
	IsDefault  uint8  `json:"is_default"`

	// 通过 customer_id 关联用户
	Customer customer.Customer `json:"customer"`

	models.CommonTimestampsField
}

func (customerContact *CustomerContact) Create() {
	database.DB.Create(&customerContact)
}

func (customerContact *CustomerContact) Save() (rowsAffected int64) {
	result := database.DB.Save(&customerContact)
	return result.RowsAffected
}

func (customerContact *CustomerContact) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&customerContact)
	return result.RowsAffected
}
