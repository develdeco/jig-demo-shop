package shipping

import "testing"

// Rate for a domestic order returns the table's 500/800/1200 cents across
// the three weight bands.
func TestRateDomestic(t *testing.T) {
	cases := []struct {
		grams int
		want  int
	}{
		{1, 500},
		{500, 500},
		{501, 800},
		{2000, 800},
		{2001, 1200},
	}
	for _, c := range cases {
		got, err := Rate("domestic", c.grams)
		if err != nil {
			t.Fatalf("Rate(domestic, %d) returned error %v, want nil", c.grams, err)
		}
		if got != c.want {
			t.Errorf("Rate(domestic, %d) = %d, want %d", c.grams, got, c.want)
		}
	}
}

// Rate for an international order returns the table's 1500/2500/4000
// cents across the three weight bands.
func TestRateInternational(t *testing.T) {
	cases := []struct {
		grams int
		want  int
	}{
		{1, 1500},
		{500, 1500},
		{501, 2500},
		{2000, 2500},
		{2001, 4000},
	}
	for _, c := range cases {
		got, err := Rate("international", c.grams)
		if err != nil {
			t.Fatalf("Rate(international, %d) returned error %v, want nil", c.grams, err)
		}
		if got != c.want {
			t.Errorf("Rate(international, %d) = %d, want %d", c.grams, got, c.want)
		}
	}
}

// Rate for an unknown region returns a non-nil error and a rate of 0,
// never a panic.
func TestRateUnknownRegion(t *testing.T) {
	got, err := Rate("mars", 100)
	if err == nil {
		t.Fatal("Rate(mars, 100) returned nil error, want non-nil")
	}
	if got != 0 {
		t.Errorf("Rate(mars, 100) = %d, want 0", got)
	}
}

// A non-positive weight falls in the lightest band, same spirit as
// cart.Total trusting non-negative prices.
func TestRateNonPositiveGrams(t *testing.T) {
	got, err := Rate("domestic", 0)
	if err != nil {
		t.Fatalf("Rate(domestic, 0) returned error %v, want nil", err)
	}
	if got != 500 {
		t.Errorf("Rate(domestic, 0) = %d, want 500", got)
	}

	got, err = Rate("domestic", -100)
	if err != nil {
		t.Fatalf("Rate(domestic, -100) returned error %v, want nil", err)
	}
	if got != 500 {
		t.Errorf("Rate(domestic, -100) = %d, want 500", got)
	}
}

// FreeShippingThresholdCents is the exported constant the brief and callers
// state their expectations against: $75.00 in cents.
func TestFreeShippingThresholdCents(t *testing.T) {
	if FreeShippingThresholdCents != 7500 {
		t.Errorf("FreeShippingThresholdCents = %d, want 7500", FreeShippingThresholdCents)
	}
}

// A subtotal at or above the threshold makes shipping free, regardless of
// region (known or not) or weight.
func TestCostAtOrAboveThresholdIsFree(t *testing.T) {
	cases := []struct {
		region        string
		grams         int
		subtotalCents int
	}{
		{"domestic", 100, FreeShippingThresholdCents},
		{"domestic", 100, FreeShippingThresholdCents + 1},
		{"international", 2001, FreeShippingThresholdCents},
		{"mars", 100, FreeShippingThresholdCents},
		{"mars", 100, FreeShippingThresholdCents + 100000},
	}
	for _, c := range cases {
		got, err := Cost(c.region, c.grams, c.subtotalCents)
		if err != nil {
			t.Errorf("Cost(%q, %d, %d) returned error %v, want nil", c.region, c.grams, c.subtotalCents, err)
		}
		if got != 0 {
			t.Errorf("Cost(%q, %d, %d) = %d, want 0", c.region, c.grams, c.subtotalCents, got)
		}
	}
}

// A subtotal below the threshold passes through to Rate unchanged, on
// success...
func TestCostBelowThresholdMatchesRate(t *testing.T) {
	cases := []struct {
		region string
		grams  int
	}{
		{"domestic", 1},
		{"domestic", 2001},
		{"international", 501},
	}
	for _, c := range cases {
		want, wantErr := Rate(c.region, c.grams)
		got, err := Cost(c.region, c.grams, FreeShippingThresholdCents-1)
		if err != wantErr {
			t.Errorf("Cost(%q, %d, threshold-1) returned error %v, want %v", c.region, c.grams, err, wantErr)
		}
		if got != want {
			t.Errorf("Cost(%q, %d, threshold-1) = %d, want %d", c.region, c.grams, got, want)
		}
	}
}

// ...and on failure: an unknown region below the threshold still returns
// Rate's non-nil error and a cost of 0, never masked by the free-shipping
// check.
func TestCostBelowThresholdUnknownRegion(t *testing.T) {
	got, err := Cost("mars", 100, FreeShippingThresholdCents-1)
	if err == nil {
		t.Fatal("Cost(mars, 100, threshold-1) returned nil error, want non-nil")
	}
	if got != 0 {
		t.Errorf("Cost(mars, 100, threshold-1) = %d, want 0", got)
	}
}
