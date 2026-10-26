// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
func Total(items []Item, discountPct int) int {
	subtotal := 0
	for _, it := range items {
		subtotal += it.PriceCents * it.Qty
	}
	// Apply the discount once, on the order, and round to the nearest cent
	// rather than truncating it away on every line.
	return (subtotal*(100-discountPct) + 50) / 100
}
