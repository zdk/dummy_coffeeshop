package main

import (
	"coffeeshop/coffee"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Coffeeshop!",
		})
	})
	r.GET("/coffee", getCoffee)
	r.Run(":8081")
}

func getCoffee(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	c.String(http.StatusOK, " %s", coffeelist)
}

// package main
//
// import (
// 	"cofeeshop/coffee"
// 	"fmt"
// )
//
// func main() {
// 	// fmt.Println("Printing the lis of coffees available in the Coffee Shop")
// 	coffees, err := coffee.GetCoffees()
// 	if err != nil {
// 		fmt.Println("Error while getting coffeelist ", err)
// 		return
// 	}
//
// 	for _, element := range coffees.List {
// 		result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
// 		fmt.Println(result)
// 	}
//
// 	fmt.Println("Is Latte Available? ", coffee.IsCoffeeAvailable("Latte"))
// 	fmt.Println("Is Americano Available? ", coffee.IsCoffeeAvailable("Americano"))
// }
