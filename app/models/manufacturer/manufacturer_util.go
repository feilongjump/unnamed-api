package manufacturer

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idStr string) (manufacturer Manufacturer) {
    database.DB.Where("id", idStr).First(&manufacturer)
    return
}

func GetBy(field, value string) (manufacturer Manufacturer) {
    database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&manufacturer)
    return
}

func All() (manufacturers []Manufacturer) {
    database.DB.Find(&manufacturers)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Manufacturer{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
    return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (manufacturers []Manufacturer, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(Manufacturer{}),
		&manufacturers,
		app.V1URL(database.TableName(&Manufacturer{})),
		perPage,
	)

	return
}