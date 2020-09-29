package stats

import (
	"github.com/ahafizi/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}

	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Errorf("result is nil")
	}
}

func TestCategoriesAvg_allIsGood(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "candy", Amount: 3_000_00},
		{ID: 4, Category: "dress", Amount: 4_000_00},
		{ID: 5, Category: "dress", Amount: 4_000_00},
		{ID: 6, Category: "auto", Amount: 5_000_00},
		{ID: 7, Category: "auto", Amount: 6_000_00},
	}

	expected := map[types.Category]types.Money{
		"auto": 4_000_00,
		"food": 2_000_00,
		"candy": 3_000_00,
		"dress": 4_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_firstTest(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}

	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_SecondTest(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_ThirdTest(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_FourthTest(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}