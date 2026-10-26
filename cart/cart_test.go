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

// One item at 3.33 with 10% off is 2.997, which rounds to 3.00: a fraction of
// at least half a cent must round up, not truncate down to 2.99.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}}
	if got := Total(items, 10); got != 300 {
		t.Fatalf("Total = %d cents, want 300", got)
	}
}

// A negative discount cannot inflate the total past the subtotal: it clamps to 0% off.
func TestTotalClampsDiscountBelowZero(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total = %d cents, want 1000", got)
	}
}

// A discount over 100% cannot drive the total negative: it clamps to 100% off.
func TestTotalClampsDiscountAboveHundred(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0", got)
	}
}
