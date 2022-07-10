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
	}

	type Customer struct {
		models.BaseModel

		Name           string `gorm:"type:varchar(255);not null;index;comment:姓名;"`
		Category       string `gorm:"type:varchar(255);not null;default:'';comment:客户分类;"`
		SalesmanId     string `gorm:"type:bigint(0);not null;index;default:0;comment:业务员;"`
		MerchandiserId string `gorm:"type:bigint(0);not null;index;default:0;comment:跟单员;"`
		Grade          string `gorm:"type:tinyint(2);not null;default:0;comment:客户等级;"`
		Currency       string `gorm:"type:varchar(255);not null;default:'';comment:币种;"`
		PaymentMethod  string `gorm:"type:varchar(255);not null;default:'';comment:付款方式;"`
		Address        string `gorm:"type:varchar(255);not null;default:'';comment:客户地址;"`
		Remarks        string `gorm:"type:text;comment:备注;"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Customer{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Customer{})
	}

	migrate.Add("2022_07_09_152419_add_customer_table", up, down)
}
