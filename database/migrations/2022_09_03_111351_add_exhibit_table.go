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

	type Exhibit struct {
		models.BaseModel

		ManufacturerID uint64      `gorm:"type:bigint;not null;index;comment:厂家;"`
		Mo             string      `gorm:"type:varchar(255);not null;comment:Mo;"`
		No             string      `gorm:"type:varchar(255);not null;comment:编号;"`
		Name           models.Lang `gorm:"type:json;not null;comment:名称;"`
		Spec           string      `gorm:"type:varchar(255);not null;default:'';comment:规格;"`
		Series         string      `gorm:"type:varchar(255);not null;default:'';comment:系列;"`
		Material       string      `gorm:"type:varchar(255);not null;default:'';comment:材料;"`
		UnitPrice      float64     `gorm:"type:decimal(10,2);not null;default:0.00;comment:单价;"`
		QuotedPrice    float64     `gorm:"type:decimal(10,2);not null;default:0.00;comment:报价;"`
		TaxRebateRate  float64     `gorm:"type:decimal(10,2);not null;default:0.00;comment:退税率;"`
		Describe       models.Lang `gorm:"type:json;comment:样品描述;"`
		PackDescribe   models.Lang `gorm:"type:json;comment:包装描述;"`
		Remarks        string      `gorm:"type:text;comment:备注;"`

		// 会创建 manufacturer_id 外键的约束
		Manufacturer Manufacturer

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Exhibit{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Exhibit{})
	}

	migrate.Add("2022_09_03_111351_add_exhibit_table", up, down)
}
