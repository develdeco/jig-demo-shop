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

// An exact half-cent rounds up: 101 cents at 50% off is 50.5 cents.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	items := []Item{{PriceCents: 101, Qty: 1}}
	if got := Total(items, 50); got != 51 {
		t.Fatalf("Total = %d cents, want 51", got)
	}
}

// A discount below 0% is clamped to 0%: the order is charged in full.
func TestTotalClampsDiscountBelowZero(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}}
	if got := Total(items, -10); got != 999 {
		t.Fatalf("Total = %d cents, want 999", got)
	}
}

// A discount above 100% is clamped to 100%: the order is free.
func TestTotalClampsDiscountAboveHundred(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total = %d cents, want 0", got)
	}
}
