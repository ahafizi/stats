package main

import (
	"fmt"
	"github.com/ahafizi/bank/v2/pkg/types"
	"github.com/ahafizi/stats/pkg/stats"
)

func main() {
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

	fmt.Println(stats.Avg(payments))
}

