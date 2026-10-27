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

// A fractional cent rounds to the nearest cent rather than truncating down,
// so the fix is distinguished from plain integer division.
func TestTotalRoundsRatherThanTruncates(t *testing.T) {
	// One item at 3.33 with 10% off is 2.997, which rounds up to 3.00;
	// truncation would give 2.99.
	if got := Total([]Item{{PriceCents: 333, Qty: 1}}, 10); got != 300 {
		t.Fatalf("Total = %d cents, want 300", got)
	}

	// A one-cent item at 50% off is an exact half-cent tie, which rounds
	// up to 1 cent; truncation would give 0.
	if got := Total([]Item{{PriceCents: 1, Qty: 1}}, 50); got != 1 {
		t.Fatalf("Total = %d cents, want 1", got)
	}
}

// A discount outside 0-100 is clamped, so the order never prices above its
// full subtotal or below zero.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with negative discount = %d cents, want 1000", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with over-100 discount = %d cents, want 0", got)
	}
}
