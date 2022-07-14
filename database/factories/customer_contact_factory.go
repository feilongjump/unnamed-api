package factories

import (
	"unnamed-api/app/models/customer_contact"

	"github.com/bxcodec/faker/v3"
)

func MakeCustomerContacts(count int) []customer_contact.CustomerContact {

	var objs []customer_contact.CustomerContact

	// 设置唯一性，如 CustomerContact 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		customerContactModel := customer_contact.CustomerContact{
			CustomerID: 1,
			Name:       faker.ChineseName(),
			Phone:      faker.Phonenumber(),
			Email:      faker.Email(),
			Fax:        faker.Phonenumber(),
			IsDefault:  0,
		}
		objs = append(objs, customerContactModel)
	}

	return objs
}
