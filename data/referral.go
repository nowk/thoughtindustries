package data

type Referral struct {
	Source       string `json:"source,omitemtpy"`
	Referrer     string `json:"referrer,omitempty"`
	ReferrerType string `json:"referrerType,omitempty"`
}

/*
 * {
 *	  "source": null,
 *	  "referrer": null,
 *	  "referrerType": null
 * }
 *
 */
