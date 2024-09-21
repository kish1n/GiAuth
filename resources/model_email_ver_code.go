/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EmailVerCode struct {
	Key
	Attributes EmailVerCodeAttributes `json:"attributes"`
}
type EmailVerCodeResponse struct {
	Data     EmailVerCode `json:"data"`
	Included Included     `json:"included"`
}

type EmailVerCodeListResponse struct {
	Data     []EmailVerCode `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustEmailVerCode - returns EmailVerCode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEmailVerCode(key Key) *EmailVerCode {
	var emailVerCode EmailVerCode
	if c.tryFindEntry(key, &emailVerCode) {
		return &emailVerCode
	}
	return nil
}
