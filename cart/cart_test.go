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

func TestTotalClampsDiscountPct(t *testing.T) {
	items := []Item{{PriceCents: 1000, Qty: 1}}
	cases := []struct {
		name        string
		discountPct int
		want        int
	}{
		{"below zero clamps to no discount", -10, 1000},
		{"above 100 clamps to free", 150, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Total(items, c.discountPct); got != c.want {
				t.Fatalf("Total = %d cents, want %d", got, c.want)
			}
		})
	}
}

func TestTotalRoundsHalfCentUp(t *testing.T) {
	// Raw 101 cents at 50% off: 101*50/100 = 50.5 exactly -> rounds up to 51.
	items := []Item{{PriceCents: 101, Qty: 1}}
	if got := Total(items, 50); got != 51 {
		t.Fatalf("Total = %d cents, want 51", got)
	}
}
