package admin_user

import (
	"unnamed-api/app/models"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/hash"
)

type AdminUser struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (adminUser *AdminUser) Create() {
	database.DB.Create(&adminUser)
}

func (adminUser *AdminUser) Save() (rowsAffected int64) {
	result := database.DB.Save(&adminUser)
	return result.RowsAffected
}

func (adminUser *AdminUser) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&adminUser)
	return result.RowsAffected
}

// ComparePassword 密码是否正确
func (adminUser *AdminUser) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, adminUser.Password)
}
