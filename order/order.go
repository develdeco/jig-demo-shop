// Package order prices a whole order: its items, discount and shipping.
package order

import (
	"example.com/shop/cart"
	"example.com/shop/shipping"
)

// Total is what the order costs, in cents: cart.Total's discounted subtotal
// plus shipping.Cost for the given region and weight, checked against that
// same discounted subtotal. Once the subtotal clears
// shipping.FreeShippingThresholdCents, shipping is free; below it, an
// unknown region returns a non-nil error and a total of 0, never a panic.
//
// An unknown region errors regardless of subtotal or weight, even once the
// subtotal clears the free-shipping threshold - where shipping.Cost itself
// would treat it as free rather than error. Total checks the region with
// shipping.Rate up front for that reason, rather than relying solely on
// shipping.Cost's error.
func Total(items []cart.Item, discountPct int, region string, grams int) (int, error) {
	subtotal := cart.Total(items, discountPct)

	if _, err := shipping.Rate(region, grams); err != nil {
		return 0, err
	}

	shippingCost, err := shipping.Cost(region, grams, subtotal)
	if err != nil {
		return 0, err
	}

	return subtotal + shippingCost, nil
}
