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

// Three items at 3.34 with 10% off cost 9.018: rounded to the nearest
// cent that's 902, not 901, which is what truncating would give. This
// pins round-half-up specifically, since TestTotalRoundsOnceOnTheOrder's
// figures land on a whole cent under both rounding and truncation.
func TestTotalRoundsHalfUp(t *testing.T) {
	items := []Item{{PriceCents: 334, Qty: 1}, {PriceCents: 334, Qty: 1}, {PriceCents: 334, Qty: 1}}
	if got := Total(items, 10); got != 902 {
		t.Fatalf("Total = %d cents, want 902", got)
	}
}

// A discount outside [0, 100] clamps to that range instead of pushing the
// total negative or above the subtotal.
func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}

	if got := Total(items, -10); got != 1000 {
		t.Fatalf("Total with discountPct -10 = %d cents, want 1000", got)
	}
	if got := Total(items, 150); got != 0 {
		t.Fatalf("Total with discountPct 150 = %d cents, want 0", got)
	}
}
