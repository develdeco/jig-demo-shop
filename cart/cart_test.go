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

// A discount outside [0, 100] clamps instead of inflating the total above
// the subtotal or driving it negative.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with discount -10 = %d cents, want 1000 (clamped to 0%% off)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with discount 150 = %d cents, want 0 (clamped to 100%% off)", got)
	}
}
