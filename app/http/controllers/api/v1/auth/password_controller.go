package auth

import (
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/models/admin_user"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController) ResetByEmail(ctx *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(ctx, &request, requests.ResetByEmail); !ok {
		return
	}

	adminUserModel := admin_user.GetBy("email", request.Email)
	if adminUserModel.ID == 0 {
		response.Abort404(ctx)
	} else {
		adminUserModel.Password = request.Password
		adminUserModel.Save()

		response.Success(ctx)
	}
}
