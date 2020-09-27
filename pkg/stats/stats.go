package stats

import (
	"github.com/ahafizi/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := 0
	countOfPayments := len(payments)
	if countOfPayments == 0 {
		return types.Money(0)
	}
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += int(payment.Amount)
	}

	return types.Money(sum / countOfPayments)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}

		if category == payment.Category {
			sum += int(payment.Amount)
		}
	}


	return types.Money(sum)
}