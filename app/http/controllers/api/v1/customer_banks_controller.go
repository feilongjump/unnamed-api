package v1

import (
	"unnamed-api/app/models/customer_bank"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type CustomerBanksController struct {
	BaseAPIController
}

func (ctrl *CustomerBanksController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := customer_bank.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *CustomerBanksController) Show(ctx *gin.Context) {
	customerBankModel := customer_bank.Get(ctx.Param("customer"), ctx.Param("bank"))
	if customerBankModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, customerBankModel)
}

func (ctrl *CustomerBanksController) Store(ctx *gin.Context) {

	request := requests.CustomerBankRequest{}
	if ok := requests.Validate(ctx, &request, requests.CustomerBankSave); !ok {
		return
	}

	customerBankModel := customer_bank.CustomerBank{
		CustomerID:     request.CustomerID,
		Name:           request.Name,
		Currency:       request.Currency,
		AccountName:    request.AccountName,
		AccountNumber:  request.AccountNumber,
		AccountBank:    request.AccountBank,
		BankAddress:    request.BankAddress,
		CompanyAddress: request.CompanyAddress,
	}
	customerBankModel.Create()
	if customerBankModel.ID > 0 {
		response.Created(ctx, customerBankModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CustomerBanksController) Update(ctx *gin.Context) {

	customerBankModel := customer_bank.Get(ctx.Param("customer"), ctx.Param("bank"))
	if customerBankModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCustomerBank(ctx, customerBankModel); !ok {
	// 	response.Abort403(ctx)
	// 	return
	// }

	request := requests.CustomerBankRequest{}
	if ok := requests.Validate(ctx, &request, requests.CustomerBankSave); !ok {
		return
	}

	customerBankModel.Name = request.Name
	customerBankModel.Currency = request.Currency
	customerBankModel.AccountName = request.AccountName
	customerBankModel.AccountNumber = request.AccountNumber
	customerBankModel.AccountBank = request.AccountBank
	customerBankModel.BankAddress = request.BankAddress
	customerBankModel.CompanyAddress = request.CompanyAddress
	rowsAffected := customerBankModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, customerBankModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *CustomerBanksController) Delete(ctx *gin.Context) {

	customerBankModel := customer_bank.Get(ctx.Param("customer"), ctx.Param("bank"))
	if customerBankModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCustomerBank(ctx, customerBankModel); !ok {
	//     response.Abort403(ctx)
	//     return
	// }

	rowsAffected := customerBankModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
