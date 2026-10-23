// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
func Total(items []Item, discountPct int) int {
	total := 0
	for _, it := range items {
		total += it.PriceCents * it.Qty * (100 - discountPct) / 100
	}
	return total
}
