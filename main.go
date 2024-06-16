package main

import (
	"cofeeshop/coffee"
	"fmt"
)

func main() {
	// fmt.Println("Printing the lis of coffees available in the Coffee Shop")
	coffees, err := coffee.GetCoffees()
	if err != nil {
		fmt.Println("Error while getting coffeelist ", err)
		return
	}

	for _, element := range coffees.List {
		result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
		fmt.Println(result)
	}

	fmt.Println("Is Latte Available? ", coffee.IsCoffeeAvailable("Latte"))
	fmt.Println("Is Americano Available? ", coffee.IsCoffeeAvailable("Americano"))
}
