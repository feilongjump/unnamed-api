package requests

import (
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidateFunc func(interface{}, *gin.Context) map[string][]string

// 表单验证
//	*gin.Context
//	interface{}	验证参数
//	ValidateFunc	验证器方法作为回调函数传参
func Validate(ctx *gin.Context, request interface{}, handler ValidateFunc) bool {

	// 解析 JSON 请求
	if err := ctx.ShouldBind(request); err != nil {
		response.BadRequest(ctx, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return false
	}

	// 表单验证
	errs := handler(request, ctx)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		response.ValidationError(ctx, errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      message,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
