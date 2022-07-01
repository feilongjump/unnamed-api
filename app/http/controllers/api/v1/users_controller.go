package v1

import (
	"unnamed-api/app/models/user"
	"unnamed-api/pkg/auth"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(ctx *gin.Context) {
	userModel := auth.CurrentUser(ctx)
	response.JSON(ctx, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(ctx *gin.Context) {
	data, pager := user.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}
