package v1

import (
	"unnamed-api/app/models/manufacturer_bank"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ManufacturerBanksController struct {
	BaseAPIController
}

func (ctrl *ManufacturerBanksController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := manufacturer_bank.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *ManufacturerBanksController) Show(ctx *gin.Context) {
	manufacturerBankModel := manufacturer_bank.Get(ctx.Param("manufacturer"), ctx.Param("bank"))
	if manufacturerBankModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, manufacturerBankModel)
}

func (ctrl *ManufacturerBanksController) Store(ctx *gin.Context) {

	request := requests.ManufacturerBankRequest{}
	if ok := requests.Validate(ctx, &request, requests.ManufacturerBankSave); !ok {
		return
	}

	manufacturerBankModel := manufacturer_bank.ManufacturerBank{
		ManufacturerID: request.ManufacturerID,
		Name:           request.Name,
		Currency:       request.Currency,
		AccountName:    request.AccountName,
		AccountNumber:  request.AccountNumber,
		AccountBank:    request.AccountBank,
		BankAddress:    request.BankAddress,
		CompanyAddress: request.CompanyAddress,
	}
	manufacturerBankModel.Create()
	if manufacturerBankModel.ID > 0 {
		response.Created(ctx, manufacturerBankModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ManufacturerBanksController) Update(ctx *gin.Context) {

	manufacturerBankModel := manufacturer_bank.Get(ctx.Param("manufacturer"), ctx.Param("bank"))
	if manufacturerBankModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyManufacturerBank(ctx, manufacturerBankModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	request := requests.ManufacturerBankRequest{}
	if ok := requests.Validate(ctx, &request, requests.ManufacturerBankSave); !ok {
		return
	}

	manufacturerBankModel.Name = request.Name
	manufacturerBankModel.Currency = request.Currency
	manufacturerBankModel.AccountName = request.AccountName
	manufacturerBankModel.AccountNumber = request.AccountNumber
	manufacturerBankModel.AccountBank = request.AccountBank
	manufacturerBankModel.BankAddress = request.BankAddress
	manufacturerBankModel.CompanyAddress = request.CompanyAddress
	rowsAffected := manufacturerBankModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, manufacturerBankModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ManufacturerBanksController) Delete(ctx *gin.Context) {

	manufacturerBankModel := manufacturer_bank.Get(ctx.Param("manufacturer"), ctx.Param("bank"))
	if manufacturerBankModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyManufacturerBank(ctx, manufacturerBankModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := manufacturerBankModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
