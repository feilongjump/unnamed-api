package seeders

import (
    "fmt"
    "unnamed-api/database/factories"
    "unnamed-api/pkg/console"
    "unnamed-api/pkg/logger"
    "unnamed-api/pkg/seed"

    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedManufacturersTable", func(db *gorm.DB) {

        manufacturers  := factories.MakeManufacturers(10)

        result := db.Table("manufacturers").Create(&manufacturers)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}