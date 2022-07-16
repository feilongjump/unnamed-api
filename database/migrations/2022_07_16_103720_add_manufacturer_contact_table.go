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
	}

	type ManufacturerContact struct {
		models.BaseModel

		ManufacturerID uint64 `gorm:"type:bigint;not null;index"`
		Name           string `gorm:"type:varchar(50);not null;index;comment:姓名;"`
		Phone          string `gorm:"type:varchar(50);not null;default:'';comment:电话;"`
		Email          string `gorm:"type:varchar(50);not null;default:'';comment:e-mail;"`
		Fax            string `gorm:"type:varchar(50);not null;default:'';comment:传真;"`
		IsDefault      uint8  `gorm:"type:tinyint(1);not null;default:0;comment:是否为默认联系人;"`

		// 会创建 manufacturer_id 外键的约束
		Manufacturer Manufacturer

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ManufacturerContact{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ManufacturerContact{})
	}

	migrate.Add("2022_07_16_103720_add_manufacturer_contact_table", up, down)
}
