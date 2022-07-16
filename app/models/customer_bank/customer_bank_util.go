package customer_bank

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(customerIDStr, idStr string) (customerBank CustomerBank) {
	database.DB.
		Where("id", idStr).
		Where("customer_id", customerIDStr).
		First(&customerBank)
	return
}

func GetBy(field, value string) (customerBank CustomerBank) {
	database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&customerBank)
	return
}

func All() (customerBanks []CustomerBank) {
	database.DB.Find(&customerBanks)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(CustomerBank{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (customerBanks []CustomerBank, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.
			Where("customer_id = ?", ctx.Param("customer")).
			Model(CustomerBank{}),
		&customerBanks,
		app.V1URL(database.TableName(&CustomerBank{})),
		perPage,
	)

	return
}
