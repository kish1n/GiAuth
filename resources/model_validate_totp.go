/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ValidateTotp struct {
	Key
	Attributes ValidateTotpAttributes `json:"attributes"`
}
type ValidateTotpResponse struct {
	Data     ValidateTotp `json:"data"`
	Included Included     `json:"included"`
}

type ValidateTotpListResponse struct {
	Data     []ValidateTotp `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustValidateTotp - returns ValidateTotp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustValidateTotp(key Key) *ValidateTotp {
	var validateTotp ValidateTotp
	if c.tryFindEntry(key, &validateTotp) {
		return &validateTotp
	}
	return nil
}
