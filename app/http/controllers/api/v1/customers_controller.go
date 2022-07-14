package v1

import (
	"unnamed-api/app/models/customer"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type CustomersController struct {
	BaseAPIController
}

func (ctrl *CustomersController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := customer.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *CustomersController) Show(ctx *gin.Context) {
	customerModel := customer.Get(ctx.Param("customer"))
	if customerModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, customerModel)
}

func (ctrl *CustomersController) Store(ctx *gin.Context) {

	request := requests.CustomerRequest{}
	if ok := requests.Validate(ctx, &request, requests.CustomerSave); !ok {
		return
	}

	customerModel := customer.Customer{
		Name:           request.Name,
		Category:       request.Category,
		SalesmanId:     request.SalesmanId,
		MerchandiserId: request.MerchandiserId,
		Grade:          request.Grade,
		Currency:       request.Currency,
		PaymentMethod:  request.PaymentMethod,
		Address:        request.Address,
		Remarks:        request.Remarks,
	}
	customerModel.Create()
	if customerModel.ID > 0 {
		response.Created(ctx, customerModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CustomersController) Update(ctx *gin.Context) {

	customerModel := customer.Get(ctx.Param("customer"))
	if customerModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCustomer(ctx, customerModel); !ok {
	// 	response.Abort403(ctx)
	// 	return
	// }

	request := requests.CustomerRequest{}
	bindOk := requests.Validate(ctx, &request, requests.CustomerSave)
	if !bindOk {
		return
	}

	customerModel.Name = request.Name
	customerModel.Category = request.Category
	customerModel.SalesmanId = request.SalesmanId
	customerModel.MerchandiserId = request.MerchandiserId
	customerModel.Grade = request.Grade
	customerModel.Currency = request.Currency
	customerModel.PaymentMethod = request.PaymentMethod
	customerModel.Address = request.Address
	customerModel.Remarks = request.Remarks

	rowsAffected := customerModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, customerModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *CustomersController) Delete(ctx *gin.Context) {

	customerModel := customer.Get(ctx.Param("customer"))
	if customerModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCustomer(ctx, customerModel); !ok {
	// 	response.Abort403(ctx)
	// 	return
	// }

	rowsAffected := customerModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
