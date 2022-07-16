package migrations

import (
	"database/sql"
	"unnamed-api/app/models"
	"unnamed-api/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Manufacturer struct {
		models.BaseModel

		No          string `gorm:"type:varchar(255);not null;index;comment:厂家编号;"`
		Name        string `gorm:"type:varchar(255);not null;index;comment:厂家名称;"`
		Category    string `gorm:"type:varchar(255);not null;default:'';comment:厂家分类;"`
		PurchaserId uint64 `gorm:"type:bigint(0);not null;index;default:0;comment:采购负责人;"`
		Address     string `gorm:"type:varchar(255);not null;default:'';comment:厂家地址;"`
		Remarks     string `gorm:"type:text;comment:备注;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Manufacturer{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Manufacturer{})
	}

	migrate.Add("2022_07_16_101828_add_manufacturer_table", up, down)
}
