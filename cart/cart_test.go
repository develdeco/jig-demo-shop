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

// Three items at 3.34 with 10% off cost 9.018, which rounds up to 9.02 -
// truncation would give 9.01. This pins rounding, not truncation, at the seam.
func TestTotalRoundsUpAFraction(t *testing.T) {
	items := []Item{{PriceCents: 334, Qty: 1}, {PriceCents: 334, Qty: 1}, {PriceCents: 334, Qty: 1}}
	if got := Total(items, 10); got != 902 {
		t.Fatalf("Total = %d cents, want 902", got)
	}
}

// One item at 0.05 with 50% off is an exact half-cent tie: 2.5 cents, which
// rounds half-up (away from zero) to 3 - truncation would give 2.
func TestTotalRoundsHalfCentTieUp(t *testing.T) {
	items := []Item{{PriceCents: 5, Qty: 1}}
	if got := Total(items, 50); got != 3 {
		t.Fatalf("Total = %d cents, want 3", got)
	}
}

// A discount below 0 clamps to 0%: no discount at all.
func TestTotalClampsDiscountBelowZero(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total = %d cents, want 1000", got)
	}
}

// A discount above 100 clamps to 100%: the order is free.
func TestTotalClampsDiscountAboveHundred(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0", got)
	}
}
