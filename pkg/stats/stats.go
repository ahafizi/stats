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

	count := 0
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += int(payment.Amount)
		count++
	}

	return types.Money(sum / count)
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

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	avgByCategories := map[types.Category]types.Money{}
	countByCategories := map[types.Category]types.Money{}

	for _, payment := range payments {
		countByCategories[payment.Category] += 1
		avgByCategories[payment.Category] += payment.Amount
	}

	for key, value := range avgByCategories {
		avgByCategories[key] = value/countByCategories[key]
	}

	return avgByCategories
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	periodsDynamic := map[types.Category]types.Money{}

	if len(first) == 0 && len(second) == 0 {
		return periodsDynamic
	}

	currentPeriods := first
	if len(first) < len(second) {
		currentPeriods = second
	}

	for key := range currentPeriods {
		periodsDynamic[key] = second[key] - first[key]
	}

	return periodsDynamic
}