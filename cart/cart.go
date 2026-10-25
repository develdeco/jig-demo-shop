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
	return roundToCent(subtotal * (100 - discountPct))
}

// roundToCent takes cents-times-100 (a subtotal times a percent) down to the
// nearest cent, rounding instead of truncating.
func roundToCent(centsTimes100 int) int {
	return (centsTimes100 + 50) / 100
}
