package data

type Coupon struct {
	Code             string `json:"code,omitempty"`
	AmountOffInCents int    `json:"amountOffInCents,omitempty"`
}

/*
 * {
 *	  "code": null,
 *	  "amountOffInCents": 0
 * }
 *
 */
