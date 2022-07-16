package v1

import (
	"unnamed-api/app/models/manufacturer"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ManufacturersController struct {
	BaseAPIController
}

func (ctrl *ManufacturersController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := manufacturer.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *ManufacturersController) Show(ctx *gin.Context) {
	manufacturerModel := manufacturer.Get(ctx.Param("manufacturer"))
	if manufacturerModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, manufacturerModel)
}

func (ctrl *ManufacturersController) Store(ctx *gin.Context) {

	request := requests.ManufacturerRequest{}
	if ok := requests.Validate(ctx, &request, requests.ManufacturerSave); !ok {
		return
	}

	manufacturerModel := manufacturer.Manufacturer{
		No:          request.No,
		Name:        request.Name,
		Category:    request.Category,
		PurchaserId: request.PurchaserId,
		Address:     request.Address,
		Remarks:     request.Remarks,
	}
	manufacturerModel.Create()
	if manufacturerModel.ID > 0 {
		response.Created(ctx, manufacturerModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ManufacturersController) Update(ctx *gin.Context) {

	manufacturerModel := manufacturer.Get(ctx.Param("manufacturer"))
	if manufacturerModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyManufacturer(ctx, manufacturerModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	request := requests.ManufacturerRequest{}
	if ok := requests.Validate(ctx, &request, requests.ManufacturerSave); !ok {
		return
	}

	manufacturerModel.No = request.No
	manufacturerModel.Name = request.Name
	manufacturerModel.Category = request.Category
	manufacturerModel.PurchaserId = request.PurchaserId
	manufacturerModel.Address = request.Address
	manufacturerModel.Remarks = request.Remarks
	rowsAffected := manufacturerModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, manufacturerModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ManufacturersController) Delete(ctx *gin.Context) {

	manufacturerModel := manufacturer.Get(ctx.Param("manufacturer"))
	if manufacturerModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyManufacturer(ctx, manufacturerModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := manufacturerModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
