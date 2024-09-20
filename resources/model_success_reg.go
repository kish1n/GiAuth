/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SuccessReg struct {
	Key
	Attributes SuccessRegAttributes `json:"attributes"`
}
type SuccessRegResponse struct {
	Data     SuccessReg `json:"data"`
	Included Included   `json:"included"`
}

type SuccessRegListResponse struct {
	Data     []SuccessReg `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustSuccessReg - returns SuccessReg from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSuccessReg(key Key) *SuccessReg {
	var successReg SuccessReg
	if c.tryFindEntry(key, &successReg) {
		return &successReg
	}
	return nil
}
