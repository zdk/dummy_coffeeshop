package coffee

import "testing"

func init() {
	Coffees = CoffeeList{
		List: []CoffeeDetails{
			{"Latte", 2.5},
			{"Flat White", 2},
			{"Cappucinno", 2.25},
		},
	}
}

func TestIsCoffeeAvailable(t *testing.T) {
	type testCase struct {
		coffeeType string
		want       bool
	}

	cases := []testCase{
		{"lat", false},
		{"Latte", true},
		{"", false},
		{"cappucinno", false},
		{"Americano", false},
	}

	for _, tc := range cases {
		got := IsCoffeeAvailable(tc.coffeeType)
		if tc.want != got {
			t.Logf("Coffee Type: %s", tc.coffeeType)
			t.Errorf("Expected '%v', but got '%v' for ", tc.want, got)
		}
	}
}
