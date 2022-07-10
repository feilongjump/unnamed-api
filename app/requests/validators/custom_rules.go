package validators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
	"unnamed-api/pkg/database"

	"github.com/thedevsaddam/govalidator"
)

// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {
	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// 	not_exists:users,email 检查数据库表里是否存在同一条信息
	// 	not_exists:users,email,1 排除用户掉 id 为 1 的用户
	govalidator.AddCustomRule("not_exists", not_exists)

	// max_cn:8 中文长度设定不超过 8
	govalidator.AddCustomRule("max_cn", max_cn)

	// min_cn:2 中文长度设定不小于 8
	govalidator.AddCustomRule("min_cn", min_cn)

	// 自定义规则 existed，验证请求数据必须存在于数据库中。
	//  用于保证数据库某个字段的值是否存在，如管理员，客户，厂家。
	//  existed 参数可以有两种，一种是 1 个参数，一种是 2 个参数：
	// 	existed:users 检查数据库表里是否存在 id 为提交值的数据
	// 	existed:users,email 检查数据库表里是否存在 email 为提交值的数据
	govalidator.AddCustomRule("existed", existed)
}

func not_exists(field, rule, message string, value interface{}) error {
	rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

	// 第一个参数，表名，如：users
	tableName := rng[0]
	// 第二个参数：字段名，如：email
	dbField := rng[1]
	// 第三个参数：排除 ID，如：1
	var exceptID string
	if len(rng) > 2 {
		exceptID = rng[2]
	}

	// 用户的请求数据
	requestValue := value.(string)

	// 拼接 SQL
	query := database.DB.Table(tableName).Where(dbField+" = ?", requestValue)

	// 如果传参第三个参数，加上 SQL Where 过滤
	if len(exceptID) > 0 {
		query.Where("id != ?", exceptID)
	}

	// 查询数据库
	var count int64
	query.Count(&count)

	// 验证不通过，数据库能找到相应的数据
	if count != 0 {
		// 如果又自定义错误消息的话
		if message != "" {
			return errors.New(message)
		}

		// 默认的错误消息
		return fmt.Errorf("%v 已被占用", requestValue)
	}

	// 验证通过
	return nil
}

func existed(field, rule, message string, value interface{}) error {
	rng := strings.Split(strings.TrimPrefix(rule, "existed:"), ",")

	// 第一个参数，表名，如：users
	tableName := rng[0]
	// 第二个参数：字段名，如：email
	dbField := "id"
	if len(rng) > 1 {
		dbField = rng[2]
	}

	// 拼接 SQL
	query := database.DB.Table(tableName).Where(dbField+" = ?", value)

	// 查询数据库
	var count int64
	query.Count(&count)

	// 验证不通过，数据库能找到相应的数据
	if count == 0 {
		// 如果又自定义错误消息的话
		if message != "" {
			return errors.New(message)
		}

		// 默认的错误消息
		return fmt.Errorf("%v 不存在", value)
	}

	// 验证通过
	return nil
}

func max_cn(field, rule, message string, value interface{}) error {
	valLength := utf8.RuneCountInString(value.(string))

	l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
	if valLength > l {
		// 如果有自定义错误消息的话，使用自定义消息
		if message != "" {
			return errors.New(message)
		}

		return fmt.Errorf("长度不能超过 %d 个字", l)
	}

	return nil
}

func min_cn(field, rule, message string, value interface{}) error {
	valLength := utf8.RuneCountInString(value.(string))

	l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
	if valLength < l {
		// 如果有自定义错误消息的话，使用自定义消息
		if message != "" {
			return errors.New(message)
		}

		return fmt.Errorf("长度需大于 %d 个字", l)
	}

	return nil
}
