package v1

import (
	"unnamed-api/app/models/exhibit"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ExhibitsController struct {
	BaseAPIController
}

func (ctrl *ExhibitsController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := exhibit.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *ExhibitsController) Show(ctx *gin.Context) {
	exhibitModel := exhibit.Get(ctx.Param("exhibit"))
	if exhibitModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, exhibitModel)
}

func (ctrl *ExhibitsController) Store(ctx *gin.Context) {

	request := requests.ExhibitRequest{}
	if ok := requests.Validate(ctx, &request, requests.ExhibitSave); !ok {
		return
	}

	exhibitModel := exhibit.Exhibit{
		ManufacturerID: request.ManufacturerID,
		Mo:             request.Mo,
		No:             request.No,
		Name:           request.Name,
		Spec:           request.Spec,
		Series:         request.Series,
		Material:       request.Material,
		UnitPrice:      request.UnitPrice,
		QuotedPrice:    request.QuotedPrice,
		TaxRebateRate:  request.TaxRebateRate,
		Describe:       request.Describe,
		PackDescribe:   request.PackDescribe,
		Remarks:        request.Remarks,
	}
	exhibitModel.Create()
	if exhibitModel.ID > 0 {
		response.Created(ctx, exhibitModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ExhibitsController) Update(ctx *gin.Context) {

	exhibitModel := exhibit.Get(ctx.Param("exhibit"))
	if exhibitModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyExhibit(ctx, exhibitModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	request := requests.ExhibitRequest{}
	if ok := requests.Validate(ctx, &request, requests.ExhibitSave); !ok {
		return
	}

	exhibitModel.Mo = request.Mo
	exhibitModel.No = request.No
	exhibitModel.Name = request.Name
	exhibitModel.Spec = request.Spec
	exhibitModel.Series = request.Series
	exhibitModel.Material = request.Material
	exhibitModel.UnitPrice = request.UnitPrice
	exhibitModel.QuotedPrice = request.QuotedPrice
	exhibitModel.TaxRebateRate = request.TaxRebateRate
	exhibitModel.Describe = request.Describe
	exhibitModel.PackDescribe = request.PackDescribe
	exhibitModel.Remarks = request.Remarks
	rowsAffected := exhibitModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, exhibitModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ExhibitsController) Delete(ctx *gin.Context) {

	exhibitModel := exhibit.Get(ctx.Param("exhibit"))
	if exhibitModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyExhibit(ctx, exhibitModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := exhibitModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
