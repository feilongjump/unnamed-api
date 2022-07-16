package manufacturer

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/admin_user"
	"unnamed-api/pkg/database"
)

type Manufacturer struct {
	models.BaseModel

	No          string `json:"no,omitempty"`
	Name        string `json:"name,omitempty"`
	Category    string `json:"category"`
	PurchaserId uint64 `json:"purchaser_id"`
	Address     string `json:"address"`
	Remarks     string `json:"remarks"`

	// 通过 purchaser_id 关联管理员
	Purchaser admin_user.AdminUser `json:"purchaser"`

	models.CommonTimestampsField
}

func (manufacturer *Manufacturer) Create() {
	database.DB.Create(&manufacturer)
}

func (manufacturer *Manufacturer) Save() (rowsAffected int64) {
	result := database.DB.Save(&manufacturer)
	return result.RowsAffected
}

func (manufacturer *Manufacturer) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&manufacturer)
	return result.RowsAffected
}
