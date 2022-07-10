package factories

import (
	"unnamed-api/app/models/customer"

	"github.com/bxcodec/faker/v3"
)

func MakeCustomers(count int) []customer.Customer {

	var objs []customer.Customer

	// 设置唯一性，如 Customer 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		customerModel := customer.Customer{
			Name:           faker.ChineseName(),
			Category:       faker.Word(),
			SalesmanId:     1,
			MerchandiserId: 1,
			Grade:          1,
			Currency:       faker.Currency(),
			PaymentMethod:  faker.CCType(),
			Address:        faker.Sentence(),
			Remarks:        faker.Paragraph(),
		}
		objs = append(objs, customerModel)
	}

	return objs
}
