/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserPublic struct {
	Key
}
type UserPublicResponse struct {
	Data     UserPublic `json:"data"`
	Included Included   `json:"included"`
}

type UserPublicListResponse struct {
	Data     []UserPublic `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustUserPublic - returns UserPublic from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserPublic(key Key) *UserPublic {
	var userPublic UserPublic
	if c.tryFindEntry(key, &userPublic) {
		return &userPublic
	}
	return nil
}
