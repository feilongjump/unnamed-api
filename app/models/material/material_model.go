package material

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/manufacturer"
	"unnamed-api/pkg/database"
)

type Material struct {
	models.BaseModel

	ManufacturerID uint64  `json:"manufacturer_id,omitempty"`
	No             string  `json:"no,omitempty"`
	Name           string  `json:"name,omitempty"`
	Spec           string  `json:"spec"`
	Category       string  `json:"category"`
	Unit           string  `json:"unit"`
	UnitPrice      float64 `json:"unit_price"`
	Remarks        string  `json:"remarks"`

	// 通过 manufacturer_id 关联厂家
	Manufacturer manufacturer.Manufacturer `json:"manufacturer"`

	models.CommonTimestampsField
}

func (material *Material) Create() {
	database.DB.Create(&material)
}

func (material *Material) Save() (rowsAffected int64) {
	result := database.DB.Save(&material)
	return result.RowsAffected
}

func (material *Material) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&material)
	return result.RowsAffected
}
