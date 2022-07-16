package factories

import (
	"unnamed-api/app/models/manufacturer_contact"

	"github.com/bxcodec/faker/v3"
)

func MakeManufacturerContacts(count int) []manufacturer_contact.ManufacturerContact {

	var objs []manufacturer_contact.ManufacturerContact

	// 设置唯一性，如 ManufacturerContact 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		manufacturerContactModel := manufacturer_contact.ManufacturerContact{
			ManufacturerID: 1,
			Name:           faker.ChineseName(),
			Phone:          faker.Phonenumber(),
			Email:          faker.Email(),
			Fax:            faker.Phonenumber(),
			IsDefault:      0,
		}
		objs = append(objs, manufacturerContactModel)
	}

	return objs
}
