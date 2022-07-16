package manufacturer_bank

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(manufacturerIDStr, idStr string) (manufacturerBank ManufacturerBank) {
	database.DB.
		Where("id", idStr).
		Where("manufacturer_id", manufacturerIDStr).
		First(&manufacturerBank)
	return
}

func GetBy(field, value string) (manufacturerBank ManufacturerBank) {
	database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&manufacturerBank)
	return
}

func All() (manufacturerBanks []ManufacturerBank) {
	database.DB.Find(&manufacturerBanks)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(ManufacturerBank{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (manufacturerBanks []ManufacturerBank, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.
			Where("manufacturer_id = ?", ctx.Param("manufacturer")).
			Model(ManufacturerBank{}),
		&manufacturerBanks,
		app.V1URL(database.TableName(&ManufacturerBank{})),
		perPage,
	)

	return
}
