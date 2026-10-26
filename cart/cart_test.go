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

// A discount below 0 or above 100 is clamped into [0, 100], so it can
// never produce a negative total or a total above the subtotal.
func TestTotalClampsOutOfRangeDiscount(t *testing.T) {
	items := []Item{{PriceCents: 500, Qty: 2}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with -10%% = %d cents, want 1000 (clamped to 0%%)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% = %d cents, want 0 (clamped to 100%%)", got)
	}
}

// A subtotal of 335 at 10% off is 301.5 exactly: truncation would give 301,
// so this pins rounding to the nearest cent rather than truncating.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	items := []Item{{PriceCents: 335, Qty: 1}}
	if got := Total(items, 10); got != 302 {
		t.Fatalf("Total = %d cents, want 302", got)
	}
}

// A credit or refund line can carry a negative PriceCents, giving a
// negative subtotal. The half-cent bias must still round away from zero,
// not toward it: -100 at 25% off is exactly -75.00, and -335 at 10% off is
// -301.5, which rounds to -302.
func TestTotalRoundsNegativeSubtotalAwayFromZero(t *testing.T) {
	if got := Total([]Item{{PriceCents: -100, Qty: 1}}, 25); got != -75 {
		t.Fatalf("Total = %d cents, want -75", got)
	}
	if got := Total([]Item{{PriceCents: -335, Qty: 1}}, 10); got != -302 {
		t.Fatalf("Total = %d cents, want -302", got)
	}
}
