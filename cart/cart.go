// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
// The discount is taken once on the order's subtotal and rounded to the
// nearest cent, rather than per line. discountPct is clamped to [0, 100].
func Total(items []Item, discountPct int) int {
	subtotal := 0
	for _, it := range items {
		subtotal += it.PriceCents * it.Qty
	}
	switch {
	case discountPct < 0:
		discountPct = 0
	case discountPct > 100:
		discountPct = 100
	}
	return (subtotal*(100-discountPct) + 50) / 100
}
