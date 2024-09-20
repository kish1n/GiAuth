/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SuccessAuth struct {
	Key
	Attributes SuccessAuthAttributes `json:"attributes"`
}
type SuccessAuthResponse struct {
	Data     SuccessAuth `json:"data"`
	Included Included    `json:"included"`
}

type SuccessAuthListResponse struct {
	Data     []SuccessAuth `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustSuccessAuth - returns SuccessAuth from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSuccessAuth(key Key) *SuccessAuth {
	var successAuth SuccessAuth
	if c.tryFindEntry(key, &successAuth) {
		return &successAuth
	}
	return nil
}
