package pricing

import "testing"

func TestCalculateTotalCents(t *testing.T) {
	tests := []struct {
		name     string
		order    Order
		expected int
	}{
		{name: "regular_it_baseline", order: Order{CustomerType: "regular", SubtotalCents: 10000, Country: "IT", CouponCode: "", BlackFriday: false}, expected: 12900},
		{name: "new_de_baseline", order: Order{CustomerType: "new", SubtotalCents: 10000, Country: "DE", CouponCode: "", BlackFriday: false}, expected: 12800},
		{name: "unknown_customer_other_country", order: Order{CustomerType: "guest", SubtotalCents: 10000, Country: "FR", CouponCode: "", BlackFriday: false}, expected: 12500},
		{name: "premium_below_10000", order: Order{CustomerType: "premium", SubtotalCents: 9999, Country: "US", CouponCode: "", BlackFriday: false}, expected: 11663},
		{name: "premium_at_10000", order: Order{CustomerType: "premium", SubtotalCents: 10000, Country: "DE", CouponCode: "", BlackFriday: false}, expected: 11610},
		{name: "save10_at_threshold", order: Order{CustomerType: "regular", SubtotalCents: 5000, Country: "DE", CouponCode: "SAVE10", BlackFriday: false}, expected: 6255},
		{name: "save10_below_threshold", order: Order{CustomerType: "regular", SubtotalCents: 4999, Country: "DE", CouponCode: "SAVE10", BlackFriday: false}, expected: 6848},
		{name: "viponly_true_with_trim", order: Order{CustomerType: " vip ", SubtotalCents: 18000, Country: " IT ", CouponCode: " VIPONLY ", BlackFriday: false}, expected: 17980},
		{name: "viponly_false_for_non_vip", order: Order{CustomerType: "premium", SubtotalCents: 10000, Country: "DE", CouponCode: "VIPONLY", BlackFriday: false}, expected: 11610},
		{name: "bulk_at_threshold", order: Order{CustomerType: "regular", SubtotalCents: 20000, Country: "FR", CouponCode: "BULK", BlackFriday: false}, expected: 21100},
		{name: "bulk_below_threshold", order: Order{CustomerType: "regular", SubtotalCents: 19999, Country: "FR", CouponCode: "BULK", BlackFriday: false}, expected: 22499},
		{name: "black_friday_regular_us", order: Order{CustomerType: "regular", SubtotalCents: 10000, Country: "US", CouponCode: "", BlackFriday: true}, expected: 11965},
		{name: "black_friday_employee_us", order: Order{CustomerType: "employee", SubtotalCents: 10000, Country: "US", CouponCode: "", BlackFriday: true}, expected: 9790},
		{name: "black_friday_non_us_no_shipping_surcharge", order: Order{CustomerType: "regular", SubtotalCents: 10000, Country: "DE", CouponCode: "", BlackFriday: true}, expected: 12205},
		{name: "freeship_at_threshold", order: Order{CustomerType: "regular", SubtotalCents: 8000, Country: "DE", CouponCode: "FREESHIP", BlackFriday: false}, expected: 9520},
		{name: "freeship_below_threshold", order: Order{CustomerType: "regular", SubtotalCents: 7999, Country: "DE", CouponCode: "FREESHIP", BlackFriday: false}, expected: 10418},
		{name: "freeship_overrides_black_friday_us_surcharge", order: Order{CustomerType: "regular", SubtotalCents: 9000, Country: "US", CouponCode: "FREESHIP", BlackFriday: true}, expected: 9148},
		{name: "vip_free_shipping_exact_threshold", order: Order{CustomerType: "vip", SubtotalCents: 17648, Country: "DE", CouponCode: "", BlackFriday: false}, expected: 17850},
		{name: "premium_free_shipping_exact_threshold", order: Order{CustomerType: "premium", SubtotalCents: 22223, Country: "DE", CouponCode: "", BlackFriday: false}, expected: 23800},
		{name: "premium_free_shipping_just_below", order: Order{CustomerType: "premium", SubtotalCents: 22222, Country: "DE", CouponCode: "", BlackFriday: false}, expected: 24698},
		{name: "taxfree_non_it", order: Order{CustomerType: "regular", SubtotalCents: 10000, Country: "DE", CouponCode: "TAXFREE", BlackFriday: false}, expected: 10900},
		{name: "taxfree_it_has_no_effect", order: Order{CustomerType: "regular", SubtotalCents: 10000, Country: "IT", CouponCode: "TAXFREE", BlackFriday: false}, expected: 12900},
		{name: "vip_it_tax_override", order: Order{CustomerType: "vip", SubtotalCents: 10000, Country: "IT", CouponCode: "", BlackFriday: false}, expected: 10900},
		{name: "tax_us_standard_rate", order: Order{CustomerType: "regular", SubtotalCents: 10000, Country: "US", CouponCode: "", BlackFriday: false}, expected: 12200},
		{name: "employee_it_no_shipping_surcharge", order: Order{CustomerType: "employee", SubtotalCents: 10000, Country: "IT", CouponCode: "", BlackFriday: false}, expected: 9240},
		{name: "employee_save10_hits_40_percent_boundary", order: Order{CustomerType: "employee", SubtotalCents: 5000, Country: "DE", CouponCode: "SAVE10", BlackFriday: false}, expected: 4970},
		{name: "employee_freeship_still_pays_surcharge", order: Order{CustomerType: "employee", SubtotalCents: 12000, Country: "DE", CouponCode: "FREESHIP", BlackFriday: false}, expected: 10496},
		{name: "negative_total_clamped_to_zero", order: Order{CustomerType: "regular", SubtotalCents: -10000, Country: "FR", CouponCode: "", BlackFriday: false}, expected: 0},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if got := CalculateTotalCents(testCase.order); got != testCase.expected {
				t.Fatalf("got %d, want %d", got, testCase.expected)
			}
		})
	}
}