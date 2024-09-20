/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserFormRegAttributes struct {
	// User birthday
	Birthday string `json:"birthday"`
	// User email
	Email string `json:"email"`
	// User name
	FirstName string `json:"first_name"`
	// User surname or if user haven`t surname it`s can be patronymic
	LastName string `json:"last_name"`
	// Can be null, patronymic, middle name, mother's surname, else
	MiddleName string `json:"middle_name"`
	// Encrypted user password
	Password string `json:"password"`
}
