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

	type ManufacturerBank struct {
		models.BaseModel

		ManufacturerID uint64 `gorm:"type:bigint;not null;index"`
		Name           string `gorm:"type:varchar(50);not null;comment:账户名称;"`
		Currency       string `gorm:"type:varchar(50);not null;default:'';comment:币种;"`
		AccountName    string `gorm:"type:varchar(50);not null;default:'';comment:开户名;"`
		AccountNumber  string `gorm:"type:varchar(100);not null;default:'';comment:银行账号;"`
		AccountBank    string `gorm:"type:varchar(100);not null;default:'';comment:所属银行;"`
		BankAddress    string `gorm:"type:varchar(255);not null;default:'';comment:银行地址;"`
		CompanyAddress string `gorm:"type:varchar(255);not null;default:'';comment:公司地址;"`

		// 会创建 manufacturer_id 外键的约束
		Manufacturer Manufacturer

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ManufacturerBank{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ManufacturerBank{})
	}

	migrate.Add("2022_07_16_104802_add_manufacturer_bank_table", up, down)
}
