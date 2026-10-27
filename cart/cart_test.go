package cart

import "testing"

// Three items at 3.33 with 10% off cost 8.991, so 8.99: the discount is
// taken once, on the order, and rounded to the nearest cent.
func TestTotalRoundsOnceOnTheOrder(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}}
	if got := Total(items, 10); got != 899 {
		t.Fatalf("Total = %d cents, want 899", got)
	}
}

// A discount outside 0-100 is clamped, so the order never prices above its
// full subtotal or below zero.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with negative discount = %d cents, want 1000", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with over-100 discount = %d cents, want 0", got)
	}
}
