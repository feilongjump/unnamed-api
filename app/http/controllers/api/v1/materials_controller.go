package v1

import (
    "unnamed-api/app/models/material"
    "unnamed-api/app/requests"
    "unnamed-api/pkg/response"

    "github.com/gin-gonic/gin"
)

type MaterialsController struct {
	BaseAPIController
}

func (ctrl *MaterialsController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := material.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *MaterialsController) Show(ctx *gin.Context) {
	materialModel := material.Get(ctx.Param("material"))
	if materialModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, materialModel)
}

func (ctrl *MaterialsController) Store(ctx *gin.Context) {

	request := requests.MaterialRequest{}
	if ok := requests.Validate(ctx, &request, requests.MaterialSave); !ok {
		return
	}

	materialModel := material.Material{
		ManufacturerID: request.ManufacturerID,
		No:             request.No,
		Name:           request.Name,
		Spec:           request.Spec,
		Category:       request.Category,
		Unit:           request.Unit,
		UnitPrice:      request.UnitPrice,
		Remarks:        request.Remarks,
	}
	materialModel.Create()
	if materialModel.ID > 0 {
		response.Created(ctx, materialModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MaterialsController) Update(ctx *gin.Context) {

	materialModel := material.Get(ctx.Param("material"))
	if materialModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyMaterial(ctx, materialModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	request := requests.MaterialRequest{}
	if ok := requests.Validate(ctx, &request, requests.MaterialSave); !ok {
		return
	}

	materialModel.No = request.No
	materialModel.Name = request.Name
	materialModel.Spec = request.Spec
	materialModel.Category = request.Category
	materialModel.Unit = request.Unit
	materialModel.UnitPrice = request.UnitPrice
	materialModel.Remarks = request.Remarks
	rowsAffected := materialModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, materialModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MaterialsController) Delete(ctx *gin.Context) {

	materialModel := material.Get(ctx.Param("material"))
	if materialModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyMaterial(ctx, materialModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := materialModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
