package customer

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idStr string) (customer Customer) {
    database.DB.Where("id", idStr).First(&customer)
    return
}

func GetBy(field, value string) (customer Customer) {
    database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&customer)
    return
}

func All() (customers []Customer) {
    database.DB.Find(&customers)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Customer{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
    return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (customers []Customer, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(Customer{}),
		&customers,
		app.V1URL(database.TableName(&Customer{})),
		perPage,
	)

	return
}