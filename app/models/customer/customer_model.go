package customer

import (
	"unnamed-api/app/models"
	"unnamed-api/app/models/admin_user"
	"unnamed-api/pkg/database"
)

type Customer struct {
	models.BaseModel

	Name           string `json:"name,omitempty"`
	Category       string `json:"category"`
	SalesmanId     uint64 `json:"salesman_id"`
	MerchandiserId uint64 `json:"merchandiser_id"`
	Grade          int8   `json:"grade"`
	Currency       string `json:"currency"`
	PaymentMethod  string `json:"payment_method"`
	Address        string `json:"address"`
	Remarks        string `json:"remarks"`

	// 通过 salesman_id 关联管理员
	Salesman admin_user.AdminUser `json:"salesman"`

	// 通过 merchandiser_id 关联管理员
	Merchandiser admin_user.AdminUser `json:"merchandiser"`

	models.CommonTimestampsField
}

func (customer *Customer) Create() {
	database.DB.Create(&customer)
}

func (customer *Customer) Save() (rowsAffected int64) {
	result := database.DB.Save(&customer)
	return result.RowsAffected
}

func (customer *Customer) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&customer)
	return result.RowsAffected
}
