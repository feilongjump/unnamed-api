package customer_contact

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(customerIDStr, idStr string) (customerContact CustomerContact) {
	database.DB.
		Where("id", idStr).
		Where("customer_id", customerIDStr).
		First(&customerContact)
	return
}

func GetBy(field, value string) (customerContact CustomerContact) {
	database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&customerContact)
	return
}

func All() (customerContacts []CustomerContact) {
	database.DB.Find(&customerContacts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(CustomerContact{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (customerContacts []CustomerContact, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.
			Where("customer_id = ?", ctx.Param("customer")).
			Model(CustomerContact{}),
		&customerContacts,
		app.V1URL(database.TableName(&CustomerContact{})),
		perPage,
	)

	return
}
