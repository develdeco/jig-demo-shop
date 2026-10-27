// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
// The discount is taken once on the order subtotal and rounded to the
// nearest cent, rather than per line. discountPct is clamped to [0, 100].
func Total(items []Item, discountPct int) int {
	discountPct = clampPct(discountPct)
	subtotal := 0
	for _, it := range items {
		subtotal += it.PriceCents * it.Qty
	}
	return roundCents(subtotal * (100 - discountPct))
}

// clampPct restricts a discount percentage to [0, 100].
func clampPct(pct int) int {
	if pct < 0 {
		return 0
	}
	if pct > 100 {
		return 100
	}
	return pct
}

// roundCents rounds hundredths of a cent to the nearest cent, half away
// from zero.
func roundCents(hundredths int) int {
	if hundredths >= 0 {
		return (hundredths + 50) / 100
	}
	return -((-hundredths + 50) / 100)
}
