// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
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
	// Apply the discount once, on the order, and round to the nearest cent
	// rather than truncating it away on every line. Go's integer division
	// truncates toward zero, so the half-cent bias has to follow the
	// numerator's sign to round away from zero on both sides.
	num := subtotal * (100 - discountPct)
	if num < 0 {
		return (num - 50) / 100
	}
	return (num + 50) / 100
}
