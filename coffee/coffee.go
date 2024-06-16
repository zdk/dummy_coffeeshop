package coffee

import (
	"fmt"

	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name  string  `mapstructure:"name"`
	Price float32 `mapstructure:"price"`
}

type CoffeeList struct {
	List []CoffeeDetails `mapstructure:"coffeelist"`
}

// Coffees is the global variable to hold the list of coffees
var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: %w", err)
		return nil, err
	}

	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(coffeetype string) bool {
	for _, element := range Coffees.List {
		if element.Name == coffeetype {
			result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
			fmt.Println(result)
			return true
		}
	}
	return false
}
