package factories

import (
	"unnamed-api/app/models/manufacturer"

	"github.com/bxcodec/faker/v3"
)

func MakeManufacturers(count int) []manufacturer.Manufacturer {

	var objs []manufacturer.Manufacturer

	// 设置唯一性，如 Manufacturer 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		manufacturerModel := manufacturer.Manufacturer{
			No:          faker.Word(),
			Name:        faker.Word(),
			Category:    faker.Word(),
			PurchaserId: 1,
			Address:     faker.Sentence(),
			Remarks:     faker.Paragraph(),
		}
		objs = append(objs, manufacturerModel)
	}

	return objs
}
