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

// A discount below 0% or above 100% clamps to the nearest boundary instead
// of producing a total outside [0, subtotal].
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}}

	if got := Total(items, -5); got != 999 {
		t.Fatalf("Total with -5%% discount = %d cents, want 999 (clamped to 0%%)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% discount = %d cents, want 0 (clamped to 100%%)", got)
	}
}
