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

// One item at 3.33 with 10% off is 2.997, so 3.00: rounded to the nearest
// cent, not truncated to 2.99.
func TestTotalRoundsToNearestCentNotDown(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}}
	if got := Total(items, 10); got != 300 {
		t.Fatalf("Total = %d cents, want 300", got)
	}
}

// A line's subtotal is its price times its quantity: two units at 5.00
// with 10% off cost 9.00, not the single-unit price discounted.
func TestTotalMultipliesLineByQty(t *testing.T) {
	items := []Item{{PriceCents: 500, Qty: 2}}
	if got := Total(items, 10); got != 900 {
		t.Fatalf("Total = %d cents, want 900", got)
	}
}

// A discount below 0% or above 100% is clamped, rather than inflating the
// total or returning a negative one.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total = %d cents, want 1000 (negative discount clamped to 0)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0 (discount over 100 clamped to 100)", got)
	}
}
