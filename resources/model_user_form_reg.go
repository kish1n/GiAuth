/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserFormReg struct {
	Key
	Attributes UserFormRegAttributes `json:"attributes"`
}
type UserFormRegRequest struct {
	Data     UserFormReg `json:"data"`
	Included Included    `json:"included"`
}

type UserFormRegListRequest struct {
	Data     []UserFormReg `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustUserFormReg - returns UserFormReg from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserFormReg(key Key) *UserFormReg {
	var userFormReg UserFormReg
	if c.tryFindEntry(key, &userFormReg) {
		return &userFormReg
	}
	return nil
}
