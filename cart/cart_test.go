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

// A half-cent result rounds up: 1 cent at 50% off is 0.5 cents, which
// truncation would drop to 0 but round-half-up takes to 1.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	items := []Item{{PriceCents: 1, Qty: 1}}
	if got := Total(items, 50); got != 1 {
		t.Fatalf("Total = %d cents, want 1", got)
	}
}

// A discount below 0% behaves as 0% off.
func TestTotalClampsNegativeDiscount(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total = %d cents, want 1000", got)
	}
}

// A discount above 100% behaves as 100% off.
func TestTotalClampsDiscountAboveFull(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0", got)
	}
}
