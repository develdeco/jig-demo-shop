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

// 111 at 10% off is 99.9, which rounds up to 100 but truncates down to 99:
// this pins rounding against a regression to per-line truncation, which
// TestTotalRoundsOnceOnTheOrder's repro cannot (999 at 10% off is 899.1,
// which floors and rounds to the same 899).
func TestTotalRoundsToNearestCent(t *testing.T) {
	items := []Item{{PriceCents: 111, Qty: 1}}
	if got := Total(items, 10); got != 100 {
		t.Fatalf("Total = %d cents, want 100", got)
	}
}

// A discount below 0% or above 100% is clamped, so it can never inflate
// the total above the subtotal or push it negative.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 500, Qty: 2}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with -10%% discount = %d cents, want 1000 (clamped to 0%%)", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% discount = %d cents, want 0 (clamped to 100%%)", got)
	}
}
