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

// A discount below 0% behaves as no discount.
func TestTotalClampsNegativeDiscountToZero(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total = %d cents, want 1000", got)
	}
}

// A discount above 100% behaves as free.
func TestTotalClampsDiscountAboveHundredToHundred(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0", got)
	}
}

// 1.01 with 1% off is 0.9999, which rounds up to 1.00: truncation would
// give 0.99.
func TestTotalRoundsUpToTheNearestCent(t *testing.T) {
	items := []Item{{PriceCents: 101, Qty: 1}}
	if got := Total(items, 1); got != 100 {
		t.Fatalf("Total = %d cents, want 100", got)
	}
}

// 0.50 with 3% off is exactly 0.485, which ties at the half cent: round-half-up
// gives 0.49, where round-half-even would give 0.48.
func TestTotalBreaksExactHalfCentTiesUp(t *testing.T) {
	items := []Item{{PriceCents: 50, Qty: 1}}
	if got := Total(items, 3); got != 49 {
		t.Fatalf("Total = %d cents, want 49", got)
	}
}
