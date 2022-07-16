package v1

import (
	"unnamed-api/app/models/manufacturer_contact"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ManufacturerContactsController struct {
	BaseAPIController
}

func (ctrl *ManufacturerContactsController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := manufacturer_contact.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *ManufacturerContactsController) Show(ctx *gin.Context) {
	manufacturerContactModel := manufacturer_contact.Get(ctx.Param("manufacturer"), ctx.Param("contact"))
	if manufacturerContactModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, manufacturerContactModel)
}

func (ctrl *ManufacturerContactsController) Store(ctx *gin.Context) {

	request := requests.ManufacturerContactRequest{}
	if ok := requests.Validate(ctx, &request, requests.ManufacturerContactSave); !ok {
		return
	}

	manufacturerContactModel := manufacturer_contact.ManufacturerContact{
		ManufacturerID: request.ManufacturerID,
		Name:           request.Name,
		Phone:          request.Phone,
		Email:          request.Email,
		Fax:            request.Fax,
		IsDefault:      request.IsDefault,
	}
	manufacturerContactModel.Create()
	if manufacturerContactModel.ID > 0 {
		response.Created(ctx, manufacturerContactModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ManufacturerContactsController) Update(ctx *gin.Context) {

	manufacturerContactModel := manufacturer_contact.Get(ctx.Param("manufacturer"), ctx.Param("contact"))
	if manufacturerContactModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyManufacturerContact(ctx, manufacturerContactModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	request := requests.ManufacturerContactRequest{}
	if ok := requests.Validate(ctx, &request, requests.ManufacturerContactSave); !ok {
		return
	}

	manufacturerContactModel.Name = request.Name
	manufacturerContactModel.Phone = request.Phone
	manufacturerContactModel.Email = request.Email
	manufacturerContactModel.Fax = request.Fax
	manufacturerContactModel.IsDefault = request.IsDefault
	rowsAffected := manufacturerContactModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, manufacturerContactModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ManufacturerContactsController) Delete(ctx *gin.Context) {

	manufacturerContactModel := manufacturer_contact.Get(ctx.Param("manufacturer"), ctx.Param("contact"))
	if manufacturerContactModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyManufacturerContact(ctx, manufacturerContactModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := manufacturerContactModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
