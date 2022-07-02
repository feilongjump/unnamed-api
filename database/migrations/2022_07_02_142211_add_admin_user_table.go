package migrations

import (
	"database/sql"
	"unnamed-api/app/models"
	"unnamed-api/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type AdminUser struct {
		models.BaseModel

		Name     string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AdminUser{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AdminUser{})
	}

	migrate.Add("2022_07_02_142211_add_admin_user_table", up, down)
}
