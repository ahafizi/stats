package stats

import (
	"fmt"
	"github.com/ahafizi/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
		},
		{
			ID: 1,
			Amount: 150_00,
			Category: "jacket",
		},
		{
			ID: 1,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 25000
}

func ExampleAvg_inFailedStatus() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
			Status: types.StatusFail,
		},
		{
			ID: 2,
			Amount: 150_00,
			Category: "jacket",
			Status: types.StatusFail,
		},
		{
			ID: 3,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 50000
}

func ExampleTotalInCategory_allIsGood() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
		},
		{
			ID: 2,
			Amount: 150_00,
			Category: "jacket",
		},
		{
			ID: 3,
			Amount: 500_00,
			Category: "dress",
		},
		{
			ID: 4,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(TotalInCategory(payments, "dress"))

	// Output:
	// 100000
}

func ExampleTotalInCategory_notInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
		},
		{
			ID: 2,
			Amount: 150_00,
			Category: "jacket",
		},
		{
			ID: 3,
			Amount: 500_00,
			Category: "dress",
		},
		{
			ID: 4,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(TotalInCategory(payments, "test"))

	// Output:
	// 0
}

func ExampleTotalInCategory_inFailedStatus() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
			Status: types.StatusFail,
		},
		{
			ID: 2,
			Amount: 150_00,
			Category: "jacket",
		},
		{
			ID: 3,
			Amount: 500_00,
			Category: "dress",
			Status: types.StatusFail,
		},
		{
			ID: 4,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(TotalInCategory(payments, "dress"))

	// Output:
	// 50000
}
