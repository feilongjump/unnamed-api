package material

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idStr string) (material Material) {
    database.DB.Where("id", idStr).First(&material)
    return
}

func GetBy(field, value string) (material Material) {
    database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&material)
    return
}

func All() (materials []Material) {
    database.DB.Find(&materials)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Material{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
    return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (materials []Material, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(Material{}),
		&materials,
		app.V1URL(database.TableName(&Material{})),
		perPage,
	)

	return
}