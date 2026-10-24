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

// A discount outside [0, 100] prices as if it had been clamped into range
// first: below 0 is no discount, above 100 is a free order.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}, {PriceCents: 333, Qty: 1}}
	if got := Total(items, -10); got != 999 {
		t.Fatalf("Total with -10%% discount = %d cents, want 999", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with 150%% discount = %d cents, want 0", got)
	}
}

// 3.34 at 15% off is a subtotal of 283.9 cents: round to the nearest cent,
// 284, not truncate to 283.
func TestTotalRoundsFractionalCentUp(t *testing.T) {
	items := []Item{{PriceCents: 334, Qty: 1}}
	if got := Total(items, 15); got != 284 {
		t.Fatalf("Total = %d cents, want 284", got)
	}
}

// A tie - exactly half a cent - rounds up, not down: 0.10 at 5% off is a
// subtotal of 9.5 cents, which rounds to 10, not truncates to 9.
func TestTotalRoundsTieUp(t *testing.T) {
	items := []Item{{PriceCents: 10, Qty: 1}}
	if got := Total(items, 5); got != 10 {
		t.Fatalf("Total = %d cents, want 10", got)
	}
}

// One line of three at 3.33 prices the same as three separate lines of one:
// qty is part of the subtotal, not just the line count.
func TestTotalMultipliesByQty(t *testing.T) {
	items := []Item{{PriceCents: 333, Qty: 3}}
	if got := Total(items, 10); got != 899 {
		t.Fatalf("Total = %d cents, want 899", got)
	}
}
