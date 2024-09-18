/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserFormAuthAttributes struct {
	// User email
	Email Email `json:"email"`
	// Can be null, patronymic, middle name, mother's surname, else
	MiddleName *string `json:"middle_name,omitempty"`
	// Encrypted user password
	Password string `json:"password"`
}
