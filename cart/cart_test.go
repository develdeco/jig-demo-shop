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

// A half-cent remainder rounds up: three items at 3.33 with 50% off is
// subtotal 999, discounted 499.5 cents, which only half-away-from-zero (not
// truncation) rounds to 500.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}}
	if got := Total(items, 50); got != 500 {
		t.Fatalf("Total = %d cents, want 500", got)
	}
}

// A discount percent outside [0, 100] is clamped before pricing: below 0
// costs the full subtotal, above 100 costs nothing.
func TestTotalClampsDiscountPercent(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with -10%% discount = %d cents, want 1000 (clamped to 0%%)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% discount = %d cents, want 0 (clamped to 100%%)", got)
	}
}
