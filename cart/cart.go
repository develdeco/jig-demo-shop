// Package cart prices an order.
package cart

// Item is one line of an order: a unit price in cents and a quantity.
type Item struct {
	PriceCents int
	Qty        int
}

// Total is what the order costs, in cents, after a percentage discount.
// The discount is taken once on the order subtotal and rounded to the
// nearest cent, rather than per line.
func Total(items []Item, discountPct int) int {
	subtotal := 0
	for _, it := range items {
		subtotal += it.PriceCents * it.Qty
	}
	return roundCents(subtotal * (100 - discountPct))
}

// roundCents rounds hundredths of a cent to the nearest cent, half away
// from zero.
func roundCents(hundredths int) int {
	if hundredths >= 0 {
		return (hundredths + 50) / 100
	}
	return -((-hundredths + 50) / 100)
}
