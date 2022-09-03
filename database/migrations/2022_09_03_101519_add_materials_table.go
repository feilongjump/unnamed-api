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

	type Material struct {
		models.BaseModel

		ManufacturerID uint64  `gorm:"type:bigint;not null;index;comment:厂家;"`
		No             string  `gorm:"type:varchar(255);not null;comment:编号;"`
		Name           string  `gorm:"type:varchar(50);not null;comment:名称;"`
		Spec           string  `gorm:"type:varchar(255);not null;default:'';comment:规格;"`
		Category       string  `gorm:"type:varchar(255);not null;default:'';comment:分类;"`
		Unit           string  `gorm:"type:varchar(100);not null;default:'';comment:单位;"`
		UnitPrice      float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:单价;"`
		Remarks        string  `gorm:"type:text;comment:备注;"`

		// 会创建 manufacturer_id 外键的约束
		Manufacturer Manufacturer

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Material{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Material{})
	}

	migrate.Add("2022_09_03_101519_add_materials_table", up, down)
}
