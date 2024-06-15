package main

import "fmt"

const (
	bubblegum     = "Bubblegum"
	toffee        = "Toffee"
	iceCream      = "Ice Cream"
	milkChocolate = "Milk Chocolate"
	doughnut      = "Doughnut"
	pancake       = "Pancake"
)

func main() {
	var earns = map[string]float32{
		bubblegum:     202,
		toffee:        118,
		iceCream:      2250,
		milkChocolate: 1680,
		doughnut:      1075,
		pancake:       80,
	}
	var income float32

	fmt.Println("Earned amount:")
	for i, v := range earns {
		fmt.Printf("%s: $%v\n", i, v)
		income += v
	}

	fmt.Printf("\nIncome: $%.1f\n", income)

	var staffExpenses float32
	fmt.Println("Staff expenses:")
	fmt.Scan(&staffExpenses)

	var otherExpenses float32
	fmt.Println("Other expenses:")
	fmt.Scan(&otherExpenses)

	fmt.Printf("Net income: $%v\n", income-staffExpenses-otherExpenses)
}
