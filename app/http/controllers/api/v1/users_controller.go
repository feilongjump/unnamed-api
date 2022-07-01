package v1

import (
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
