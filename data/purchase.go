package data

type Purchase struct {
	Id                    string     `json:"id"`
	GiftRecipientEmail    string     `json:"giftRecipientEmail,omitempty"`
	Timestamp             string     `json:"timestamp"`
	CompanyId             string     `json:"companyId"`
	Coupon                Coupon     `json:"coupon"`
	CouponCode            string     `json:"couponCode,omitempty"`
	PriceInCents          int        `json:"priceInCents"`
	PurchasableId         string     `json:"purchasableId"`
	PurchasableType       string     `json:"purchasableType"`
	ChargeId              string     `json:"chargeId"`
	GroupId               string     `json:"groupId"`
	OrderId               string     `json:"orderId,omitempty"`
	RevenueType           string     `json:"revenueType"`
	Quantity              int        `json:"quantity,omitempty"`
	Variation             string     `json:"variation,omitempty"`
	ShippingMethod        string     `json:"shippingMethod,omitempty"`
	Type                  string     `json:"type,omitempty"`
	Sku                   string     `json:"sku,omitempty"`
	Slug                  string     `json:"slug"`
	Title                 string     `json:"title"`
	CourseStartDate       string     `json:"courseStartDate"`
	AmountNet             int        `json:"amountNet,omitempty"`
	AmountWithoutDiscount int        `json:"amountWithoutDiscount,omitempty"`
	AmountDiscounted      int        `json:"amountDiscounted"`
	AmountCharged         int        `json:"amountCharged,omitempty"`
	AmountRefunded        int        `json:"amountRefunded"`
	AmountDisputed        int        `json:"amountDisputed"`
	AmountTaxed           int        `json:"amountTaxed"`
	Success               bool       `json:"success"`
	FailureCode           string     `json:"failureCode,omitempty"`
	FailureMessage        string     `json:"failureMessage,omitempty"`
	User                  string     `json:"user"`
	UserDetail            UserDetail `json:"userDetail"`
	Referral              Referral   `json:"referral"`
}

/*
 * {
 *	  "id": "66b131f7-33b7-4d5a-b0fa-bed64acd4c4e-1517868872-ch_x44IEO4smDLAidcy2xz5ZPMg-0",
 *	  "giftRecipientEmail": null,
 *	  "timestamp": "2018-02-05T22:14:32.000Z",
 *	  "companyId": "66b131f7-33b7-4d5a-b0fa-bed64acd4c4e",
 *	  "coupon": { "code": null, "amountOffInCents": 0 },
 *	  "couponCode": null,
 *	  "priceInCents": 2499,
 *	  "purchasableId": "0076e871-ecbe-480c-a315-d2592359f521",
 *	  "purchasableType": "course",
 *	  "chargeId": "ch_x44IEO4smDLAidcy2xz5ZPMg-0",
 *	  "groupId": "ch_x44IEO4smDLAidcy2xz5ZPMg",
 *	  "orderId": null,
 *	  "revenueType": "new",
 *	  "quantity": 1,
 *	  "variation": null,
 *	  "shippingMethod": null,
 *	  "type": "charge",
 *	  "sku": null,
 *	  "slug": "test-course",
 *	  "title": "Test Course",
 *	  "courseStartDate": "2017-12-18T18:03:17.073Z",
 *	  "amountNet": 2499,
 *	  "amountWithoutDiscount": 2499,
 *	  "amountDiscounted": 0,
 *	  "amountCharged": 2499,
 *	  "amountRefunded": 0,
 *	  "amountDisputed": 0,
 *	  "amountTaxed": 0,
 *	  "success": true,
 *	  "failureCode": null,
 *	  "failureMessage": null,
 *	  "user": "test@thoughtindustries.com",
 *	  "userDetail": {
 *		  "id": "1bfaca09-16b3-4f8e-88c3-878e11b0445b",
 *		  "shippingName": null,
 *		  "address1": null,
 *		  "address2": null,
 *		  "city": null,
 *		  "state": null,
 *		  "zipCode": null,
 *		  "telephone": null,
 *		  "country": null,
 *		  "firstName": "Test",
 *		  "lastName": "User",
 *		  "email": "test@thoughtindustries.com",
 *		  "stripeCustomerId": "cus_YYauXzGmWnDqSv",
 *		  "externalCustomerId": null,
 *		  "sfContactId": null,
 *		  "sfAccountId": null,
 *		  "ref1": null,
 *		  "ref2": null,
 *		  "ref3": null,
 *		  "ref4": null,
 *		  "ref5": null,
 *		  "ref6": null,
 *		  "ref7": null,
 *		  "ref8": null,
 *		  "ref9": null,
 *		  "ref10": null
 *	  },
 *	  "referral": { "source": null, "referrer": null, "referrerType": null }
 * }
 *
 */
