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

// One item at 5 cents with 90% off is 0.5 cents: an exact half-cent ties up.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	items := []Item{{PriceCents: 5, Qty: 1}}
	if got := Total(items, 90); got != 1 {
		t.Fatalf("Total = %d cents, want 1", got)
	}
}

func TestTotalClampsDiscountBelowZero(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total = %d cents, want 1000", got)
	}
}

func TestTotalClampsDiscountAboveHundred(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0", got)
	}
}
