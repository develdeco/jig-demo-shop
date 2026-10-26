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

// A fractional cent at or past the half boundary rounds up instead of
// truncating down: 55.5 cents rounds to 56, and 0.5 cents rounds to 1.
func TestTotalRoundsHalfCentUp(t *testing.T) {
	if got := Total([]Item{{PriceCents: 111, Qty: 1}}, 50); got != 56 {
		t.Fatalf("Total = %d cents, want 56", got)
	}
	if got := Total([]Item{{PriceCents: 10, Qty: 1}}, 95); got != 1 {
		t.Fatalf("Total = %d cents, want 1", got)
	}
}

// A discount below 0% or above 100% is clamped before it is applied, so it
// can never produce a negative total or one over the subtotal.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with -10%% discount = %d cents, want 1000 (clamped to 0%%)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% discount = %d cents, want 0 (clamped to 100%%)", got)
	}
}
