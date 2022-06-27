package auth

import (
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/models/user"
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

	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(ctx)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(ctx)
	}
}
