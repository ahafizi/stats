package stats

import (
	"github.com/ahafizi/bank/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := 0
	countOfPayments := len(payments)
	if countOfPayments == 0 {
		return types.Money(0)
	}
	for _, payment := range payments {
		sum += int(payment.Amount)
	}

	return types.Money(sum / countOfPayments)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if category == payment.Category {
			sum += int(payment.Amount)
		}
	}


	return types.Money(sum)
}