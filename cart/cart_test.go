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

// A single item at 7 cents with 50% off is 3.5 cents, which rounds half-up
// to 4: truncation would give 3, so this case tells the two apart.
func TestTotalRoundsHalfUp(t *testing.T) {
	items := []Item{{PriceCents: 7, Qty: 1}}
	if got := Total(items, 50); got != 4 {
		t.Fatalf("Total = %d cents, want 4 (3.5 rounded half-up)", got)
	}
}

// A line's quantity multiplies into the subtotal: one line at 3.33 x3
// prices the same as three lines at 3.33 x1.
func TestTotalMultipliesQty(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 3}}
	if got := Total(items, 10); got != 899 {
		t.Fatalf("Total = %d cents, want 899", got)
	}
}

// A discount outside [0, 100] clamps instead of discounting past free or
// into a negative price.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with -10%% = %d cents, want 1000 (clamped to 0%%)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% = %d cents, want 0 (clamped to 100%%)", got)
	}
}
