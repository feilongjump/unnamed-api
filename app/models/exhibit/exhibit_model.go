package exhibit

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/manufacturer"
	"unnamed-api/pkg/database"
)

type Exhibit struct {
	models.BaseModel

	ManufacturerID uint64      `json:"manufacturer_id"`
	Mo             string      `json:"mo"`
	No             string      `json:"no"`
	Name           models.Lang `json:"name"`
	Spec           string      `json:"spec"`
	Series         string      `json:"series"`
	Material       string      `json:"material"`
	UnitPrice      float64     `json:"unit_price"`
	QuotedPrice    float64     `json:"quoted_price"`
	TaxRebateRate  float64     `json:"tax_rebate_rate"`
	Describe       models.Lang `json:"describe"`
	PackDescribe   models.Lang `json:"pack_describe"`
	Remarks        string      `json:"remarks"`

	// 通过 manufacturer_id 关联厂家
	Manufacturer manufacturer.Manufacturer `json:"manufacturer"`

	models.CommonTimestampsField
}

func (exhibit *Exhibit) Create() {
	database.DB.Create(&exhibit)
}

func (exhibit *Exhibit) Save() (rowsAffected int64) {
	result := database.DB.Save(&exhibit)
	return result.RowsAffected
}

func (exhibit *Exhibit) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&exhibit)
	return result.RowsAffected
}
