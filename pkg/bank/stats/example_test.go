package stats

import (
	"github.com/root5427/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
		  ID: 2,
		  Amount: 1000,
		  Category: "Mobile",
		},
		{
		  ID: 1,
		  Amount: 2000,
		  Category: "Mobile",
		},
		{
		  ID: 3,
		  Amount: 3000,
		  Category: "Withdraw",
		},
	  }
	
	  fmt.Println(Avg(payments))
	
	  // Output: 
	  // 2000
}
