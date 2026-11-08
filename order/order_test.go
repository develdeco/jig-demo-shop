package order

import (
	"testing"

	"example.com/shop/cart"
	"example.com/shop/shipping"
)

// A subtotal below the free-shipping threshold returns cart.Total plus
// the rate shipping.Rate would return for that region and weight.
func TestTotalBelowThresholdAddsShippingRate(t *testing.T) {
	items := []cart.Item{{PriceCents: 1000, Qty: 1}}
	subtotal := cart.Total(items, 0)
	if subtotal >= shipping.FreeShippingThresholdCents {
		t.Fatalf("subtotal %d is not below the free-shipping threshold, fixture is wrong", subtotal)
	}

	wantRate, err := shipping.Rate("domestic", 100)
	if err != nil {
		t.Fatalf("shipping.Rate returned error %v, want nil", err)
	}

	got, err := Total(items, 0, "domestic", 100)
	if err != nil {
		t.Fatalf("Total returned error %v, want nil", err)
	}
	if want := subtotal + wantRate; got != want {
		t.Errorf("Total = %d, want %d (subtotal %d + rate %d)", got, want, subtotal, wantRate)
	}
}

// A subtotal at or above the free-shipping threshold returns exactly
// cart.Total - no shipping added.
func TestTotalAtOrAboveThresholdAddsNoShipping(t *testing.T) {
	items := []cart.Item{{PriceCents: shipping.FreeShippingThresholdCents, Qty: 1}}
	subtotal := cart.Total(items, 0)
	if subtotal < shipping.FreeShippingThresholdCents {
		t.Fatalf("subtotal %d is not at or above the free-shipping threshold, fixture is wrong", subtotal)
	}

	got, err := Total(items, 0, "domestic", 100)
	if err != nil {
		t.Fatalf("Total returned error %v, want nil", err)
	}
	if got != subtotal {
		t.Errorf("Total = %d, want %d (cart.Total alone)", got, subtotal)
	}
}

// An unknown region returns a non-nil error and a total of 0, never a
// panic, regardless of subtotal or weight - including a subtotal at or
// above the free-shipping threshold, where shipping.Cost itself would not
// error.
func TestTotalUnknownRegionErrorsRegardlessOfSubtotal(t *testing.T) {
	belowItems := []cart.Item{{PriceCents: 1000, Qty: 1}}
	aboveItems := []cart.Item{{PriceCents: shipping.FreeShippingThresholdCents, Qty: 1}}

	cases := []struct {
		name  string
		items []cart.Item
		grams int
	}{
		{"below threshold", belowItems, 100},
		{"at or above threshold", aboveItems, 100},
	}
	for _, c := range cases {
		got, err := Total(c.items, 0, "mars", c.grams)
		if err == nil {
			t.Errorf("%s: Total returned nil error, want non-nil", c.name)
		}
		if got != 0 {
			t.Errorf("%s: Total = %d, want 0", c.name, got)
		}
	}
}

// The threshold comparison uses the discounted subtotal, not the
// pre-discount sum of item prices: an order whose pre-discount price is
// above the threshold but whose discounted subtotal falls below it still
// gets shipping added.
func TestTotalThresholdUsesDiscountedSubtotal(t *testing.T) {
	items := []cart.Item{{PriceCents: shipping.FreeShippingThresholdCents + 100, Qty: 1}}
	discountPct := 50

	subtotal := cart.Total(items, discountPct)
	if subtotal >= shipping.FreeShippingThresholdCents {
		t.Fatalf("discounted subtotal %d is not below the free-shipping threshold, fixture is wrong", subtotal)
	}

	wantRate, err := shipping.Rate("domestic", 100)
	if err != nil {
		t.Fatalf("shipping.Rate returned error %v, want nil", err)
	}

	got, err := Total(items, discountPct, "domestic", 100)
	if err != nil {
		t.Fatalf("Total returned error %v, want nil", err)
	}
	if want := subtotal + wantRate; got != want {
		t.Errorf("Total = %d, want %d (discounted subtotal %d + rate %d)", got, want, subtotal, wantRate)
	}
}
