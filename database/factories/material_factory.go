package factories

import (
	"github.com/bxcodec/faker/v3"
	"unnamed-api/app/models/material"
)

func MakeMaterials(count int) []material.Material {

	var objs []material.Material

	// 设置唯一性，如 Material 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		materialModel := material.Material{
			ManufacturerID: 1,
			No:             faker.Word(),
			Name:           faker.Word(),
			Spec:           faker.Word(),
			Category:       faker.Word(),
			Unit:           faker.Currency(),
			UnitPrice:      faker.Longitude(),
			Remarks:        faker.Paragraph(),
		}
		objs = append(objs, materialModel)
	}

	return objs
}
