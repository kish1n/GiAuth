/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserFormAuth struct {
	Key
	Attributes UserFormAuthAttributes `json:"attributes"`
}
type UserFormAuthRequest struct {
	Data     UserFormAuth `json:"data"`
	Included Included     `json:"included"`
}

type UserFormAuthListRequest struct {
	Data     []UserFormAuth `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustUserFormAuth - returns UserFormAuth from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserFormAuth(key Key) *UserFormAuth {
	var userFormAuth UserFormAuth
	if c.tryFindEntry(key, &userFormAuth) {
		return &userFormAuth
	}
	return nil
}
