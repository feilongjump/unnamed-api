package migrations

import (
	"database/sql"
	"unnamed-api/app/models"
	"unnamed-api/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Customer struct {
		models.BaseModel
	}

	type CustomerContact struct {
		models.BaseModel

		CustomerID uint64 `gorm:"type:bigint;not null;index"`
		Name       string `gorm:"type:varchar(50);not null;index;comment:姓名;"`
		Phone      string `gorm:"type:varchar(50);not null;default:'';comment:电话;"`
		Email      string `gorm:"type:varchar(50);not null;default:'';comment:e-mail;"`
		Fax        string `gorm:"type:varchar(50);not null;default:'';comment:传真;"`
		IsDefault  uint8  `gorm:"type:tinyint(1);not null;default:0;comment:是否为默认联系人;"`

		// 会创建 customer_id 外键的约束
		Customer Customer

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&CustomerContact{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&CustomerContact{})
	}

	migrate.Add("2022_07_14_200420_add_customer_contact_table", up, down)
}
