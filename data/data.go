// data package is the "single" object data objects
// payload models are located in the .json directory
package data

type PageInfo struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"hasMore"`
}

/*
 * {
 *	  "cursor": "MTUxODM2NTg1Mzk5OQ",
 *	  "hasMore": true
 * }
 *
 */
