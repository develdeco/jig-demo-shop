// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
// discountPct is clamped to [0, 100] before pricing.
func Total(items []Item, discountPct int) int {
	switch {
	case discountPct < 0:
		discountPct = 0
	case discountPct > 100:
		discountPct = 100
	}

	subtotal := 0
	for _, it := range items {
		subtotal += it.PriceCents * it.Qty
	}
	return (subtotal*(100-discountPct) + 50) / 100
}
