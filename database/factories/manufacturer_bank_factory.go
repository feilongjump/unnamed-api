package factories

import (
	"unnamed-api/app/models/manufacturer_bank"

	"github.com/bxcodec/faker/v3"
)

func MakeManufacturerBanks(count int) []manufacturer_bank.ManufacturerBank {

	var objs []manufacturer_bank.ManufacturerBank

	// 设置唯一性，如 ManufacturerBank 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		manufacturerBankModel := manufacturer_bank.ManufacturerBank{
			ManufacturerID: 1,
			Name:           faker.ChineseName(),
			Currency:       faker.Currency(),
			AccountName:    faker.Name(),
			AccountNumber:  faker.CCNumber(),
			AccountBank:    faker.CCType(),
			BankAddress:    faker.Sentence(),
			CompanyAddress: faker.Sentence(),
		}
		objs = append(objs, manufacturerBankModel)
	}

	return objs
}
