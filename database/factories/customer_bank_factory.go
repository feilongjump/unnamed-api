package factories

import (
	"unnamed-api/app/models/customer_bank"

	"github.com/bxcodec/faker/v3"
)

func MakeCustomerBanks(count int) []customer_bank.CustomerBank {

	var objs []customer_bank.CustomerBank

	// 设置唯一性，如 CustomerBank 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		customerBankModel := customer_bank.CustomerBank{
			CustomerID:     1,
			Name:           faker.ChineseName(),
			Currency:       faker.Currency(),
			AccountName:    faker.Name(),
			AccountNumber:  faker.CCNumber(),
			AccountBank:    faker.CCType(),
			BankAddress:    faker.Sentence(),
			CompanyAddress: faker.Sentence(),
		}
		objs = append(objs, customerBankModel)
	}

	return objs
}
