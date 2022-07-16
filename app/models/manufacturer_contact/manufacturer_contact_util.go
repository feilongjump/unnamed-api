package manufacturer_contact

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(manufacturerIDStr, idStr string) (manufacturerContact ManufacturerContact) {
	database.DB.
		Where("id", idStr).
		Where("manufacturer_id", manufacturerIDStr).
		First(&manufacturerContact)
	return
}

func GetBy(field, value string) (manufacturerContact ManufacturerContact) {
	database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&manufacturerContact)
	return
}

func All() (manufacturerContacts []ManufacturerContact) {
	database.DB.Find(&manufacturerContacts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(ManufacturerContact{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (manufacturerContacts []ManufacturerContact, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.
			Where("manufacturer_id = ?", ctx.Param("manufacturer")).
			Model(ManufacturerContact{}),
		&manufacturerContacts,
		app.V1URL(database.TableName(&ManufacturerContact{})),
		perPage,
	)

	return
}
