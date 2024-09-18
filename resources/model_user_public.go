/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "net/mail"

type UserPublic struct {
	// User email
	Email mail.Address `json:"email"`
	// User name
	FirstName string `json:"first_name"`
	// User surname or if user haven`t surname it`s can be patronymic
	LastName string `json:"last_name"`
	// Can be null, patronymic, middle name, mother's surname, else
	MiddleName string `json:"middle_name"`
	// unique user identifier is available to all
	Username string `json:"username"`
}
