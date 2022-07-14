package v1

import (
	"unnamed-api/app/models/customer_contact"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type CustomerContactsController struct {
	BaseAPIController
}

func (ctrl *CustomerContactsController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := customer_contact.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *CustomerContactsController) Show(ctx *gin.Context) {
	customerContactModel := customer_contact.Get(ctx.Param("customer"), ctx.Param("contact"))
	if customerContactModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.JSON(ctx, customerContactModel)
}

func (ctrl *CustomerContactsController) Store(ctx *gin.Context) {

	request := requests.CustomerContactRequest{}
	if ok := requests.Validate(ctx, &request, requests.CustomerContactSave); !ok {
		return
	}

	customerContactModel := customer_contact.CustomerContact{
		CustomerID: request.CustomerID,
		Name:       request.Name,
		Phone:      request.Phone,
		Email:      request.Email,
		Fax:        request.Fax,
		IsDefault:  request.IsDefault,
	}
	customerContactModel.Create()
	if customerContactModel.ID > 0 {
		response.Created(ctx, customerContactModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CustomerContactsController) Update(ctx *gin.Context) {

	customerContactModel := customer_contact.Get(ctx.Param("customer"), ctx.Param("contact"))
	if customerContactModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCustomerContact(ctx, customerContactModel); !ok {
	// 	response.Abort403(ctx)
	// 	return
	// }

	request := requests.CustomerContactRequest{}
	if ok := requests.Validate(ctx, &request, requests.CustomerContactSave); !ok {
		return
	}

	customerContactModel.Name = request.Name
	customerContactModel.Phone = request.Phone
	customerContactModel.Email = request.Email
	customerContactModel.Fax = request.Fax
	customerContactModel.IsDefault = request.IsDefault
	rowsAffected := customerContactModel.Save()
	if rowsAffected > 0 {
		response.JSON(ctx, customerContactModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}

func (ctrl *CustomerContactsController) Delete(ctx *gin.Context) {

	customerContactModel := customer_contact.Get(ctx.Param("customer"), ctx.Param("contact"))
	if customerContactModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// if ok := policies.CanModifyCustomerContact(ctx, customerContactModel); !ok {
	// 	response.Abort403(ctx)
	// 	return
	// }

	rowsAffected := customerContactModel.Delete()
	if rowsAffected > 0 {
		response.Success(ctx)
		return
	}

	response.Abort500(ctx, "删除失败，请稍后尝试~")
}
