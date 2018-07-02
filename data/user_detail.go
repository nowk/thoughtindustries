package data

type UserDetail struct {
	Id                 string `json:"id"`
	ShippingName       string `json:"shippingName,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Address2           string `json:"address2,omitempty"`
	City               string `json:"city,omitempty"`
	State              string `json:"state,omitempty"`
	ZipCode            string `json:"zipCode,omitempty"`
	Telephone          string `json:"telephone,omitempty"`
	Country            string `json:"country,omitempty"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Email              string `json:"email"`
	StripeCustomerId   string `json:"stripeCustomerId,omitempty"`
	ExternalCustomerId string `json:"externalCustomerId,omitempty"`
	SfContactId        string `json:"sfContactId,omitempty"`
	SfAccountId        string `json:"sfAccountId,omitempty"`
	Ref1               string `json:"ref1,omitempty"`
	Ref2               string `json:"ref2,omitempty"`
	Ref3               string `json:"ref3,omitempty"`
	Ref4               string `json:"ref4,omitempty"`
	Ref5               string `json:"ref5,omitempty"`
	Ref6               string `json:"ref6,omitempty"`
	Ref7               string `json:"ref7,omitempty"`
	Ref8               string `json:"ref8,omitempty"`
	Ref9               string `json:"ref9,omitempty"`
	Ref10              string `json:"ref10,omitempty"`
}

/*
 * {
 *   "id": "1bfaca09-16b3-4f8e-88c3-878e11b0445b",
 *   "shippingName": null,
 *   "address1": null,
 *   "address2": null,
 *   "city": null,
 *   "state": null,
 *   "zipCode": null,
 *   "telephone": null,
 *   "country": null,
 *   "firstName": "Test",
 *   "lastName": "User",
 *   "email": "test@thoughtindustries.com",
 *   "stripeCustomerId": "cus_YYauXzGmWnDqSv",
 *   "externalCustomerId": null,
 *   "sfContactId": null,
 *   "sfAccountId": null,
 *   "ref1": null,
 *   "ref2": null,
 *   "ref3": null,
 *   "ref4": null,
 *   "ref5": null,
 *   "ref6": null,
 *   "ref7": null,
 *   "ref8": null,
 *   "ref9": null,
 *   "ref10": null
 * }
 *
 */
