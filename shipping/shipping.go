// Package shipping rates an order by region and weight.
package shipping

import "fmt"

// rates is the shipping cost in cents, keyed by region, for the three
// weight bands in order: <=500g, 500-2000g, >2000g.
var rates = map[string][3]int{
	"domestic":      {500, 800, 1200},
	"international": {1500, 2500, 4000},
}

// Rate is the shipping cost in cents for a region and a weight in grams,
// looked up in a table of weight bands: <=500g, 500-2000g, and >2000g.
// Rate trusts its caller for a sane, non-negative weight; a non-positive
// grams falls in the lightest band. An unknown region returns a rate of
// 0 and a non-nil error, never a panic.
func Rate(region string, grams int) (int, error) {
	bands, ok := rates[region]
	if !ok {
		return 0, fmt.Errorf("shipping: unknown region %q", region)
	}
	switch {
	case grams <= 500:
		return bands[0], nil
	case grams <= 2000:
		return bands[1], nil
	default:
		return bands[2], nil
	}
}
