package v1

import (
	"unnamed-api/app/models/user"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/auth"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminUsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录管理员用户信息
func (ctrl *AdminUsersController) CurrentAdminUser(ctx *gin.Context) {
	adminUserModel := auth.CurrentAdminUser(ctx)
	response.JSON(ctx, adminUserModel)
}

// Index 所有管理员用户
func (ctrl *AdminUsersController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}
