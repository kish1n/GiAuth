/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SuccesLogout struct {
	Key
	Attributes SuccesLogoutAttributes `json:"attributes"`
}
type SuccesLogoutResponse struct {
	Data     SuccesLogout `json:"data"`
	Included Included     `json:"included"`
}

type SuccesLogoutListResponse struct {
	Data     []SuccesLogout `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustSuccesLogout - returns SuccesLogout from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSuccesLogout(key Key) *SuccesLogout {
	var succesLogout SuccesLogout
	if c.tryFindEntry(key, &succesLogout) {
		return &succesLogout
	}
	return nil
}
