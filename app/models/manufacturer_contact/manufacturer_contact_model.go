package manufacturer_contact

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/manufacturer"
	"unnamed-api/pkg/database"
)

type ManufacturerContact struct {
	models.BaseModel

	ManufacturerID uint64 `json:"manufacturer_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Fax            string `json:"fax"`
	IsDefault      uint8  `json:"is_default"`

	// 通过 manufacturer_id 关联用户
	Manufacturer manufacturer.Manufacturer `json:"manufacturer"`

	models.CommonTimestampsField
}

func (manufacturerContact *ManufacturerContact) Create() {
	database.DB.Create(&manufacturerContact)
}

func (manufacturerContact *ManufacturerContact) Save() (rowsAffected int64) {
	result := database.DB.Save(&manufacturerContact)
	return result.RowsAffected
}

func (manufacturerContact *ManufacturerContact) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&manufacturerContact)
	return result.RowsAffected
}
