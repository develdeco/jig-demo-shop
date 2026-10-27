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

// One item at 1.01 with 50% off is 50.5 cents, a half-cent tie: only
// round-half-up gives 51. Truncation and round-half-down both give 50,
// so this is the case that pins the tie direction the brief names.
func TestTotalRoundsHalfUp(t *testing.T) {
	items := []Item{{PriceCents: 101, Qty: 1}}
	if got := Total(items, 50); got != 51 {
		t.Fatalf("Total = %d cents, want 51", got)
	}
}

// A discount outside [0, 100] clamps to that range instead of pushing the
// total negative or above the subtotal.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with discountPct -10 = %d cents, want 1000", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with discountPct 150 = %d cents, want 0", got)
	}
}
