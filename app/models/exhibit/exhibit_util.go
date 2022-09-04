package exhibit

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idStr string) (exhibit Exhibit) {
    database.DB.Where("id", idStr).First(&exhibit)
    return
}

func GetBy(field, value string) (exhibit Exhibit) {
    database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&exhibit)
    return
}

func All() (exhibits []Exhibit) {
    database.DB.Find(&exhibits)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Exhibit{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
    return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (exhibits []Exhibit, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(Exhibit{}),
		&exhibits,
		app.V1URL(database.TableName(&Exhibit{})),
		perPage,
	)

	return
}