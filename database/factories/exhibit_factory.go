package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/spf13/cast"
	"strings"
	"unnamed-api/app/models"
	"unnamed-api/app/models/exhibit"
)

func MakeExhibits(count int) []exhibit.Exhibit {

	var objs []exhibit.Exhibit

	// 设置唯一性，如 Exhibit 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {

		amount := faker.AmountWithCurrency()
		amountIndex := strings.Index(amount, " ")
		unitPrice := cast.ToFloat64(amount[amountIndex+1:])

		amount = faker.AmountWithCurrency()
		amountIndex = strings.Index(amount, " ")
		quotedPrice := cast.ToFloat64(amount[amountIndex+1:])

		amount = faker.AmountWithCurrency()
		amountIndex = strings.Index(amount, " ")
		TaxRebateRate := cast.ToFloat64(amount[amountIndex+1:])

		exhibitModel := exhibit.Exhibit{
			ManufacturerID: 1,
			Mo:             faker.Word(),
			No:             faker.Word(),
			Name: models.Lang{
				Zh: faker.Word(),
			},
			Spec:          faker.Word(),
			Series:        faker.Word(),
			Material:      faker.Word(),
			UnitPrice:     unitPrice,
			QuotedPrice:   quotedPrice,
			TaxRebateRate: TaxRebateRate,
			Describe: models.Lang{
				Zh: faker.Sentence(),
				En: faker.Sentence(),
			},
			PackDescribe: models.Lang{
				Zh: faker.Sentence(),
				En: faker.Sentence(),
			},
			Remarks: faker.Paragraph(),
		}
		objs = append(objs, exhibitModel)
	}

	return objs
}
