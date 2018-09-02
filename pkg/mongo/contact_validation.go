package mongo

/** Contact model validation **/

import (
	"errors"
	"mgo-n-goji/pkg"
	"regexp"
)

//isValidContact verifies if a contact has a valid nick and at least one valid mobile. Mobile and email lists are also verified too.
func isValidContact(contact *root.Contact) (bool, error) {
	if contact.Nick == "" {
		return false, errors.New("Nick can not be empty")
	}
	if !isValidNick(contact.Nick) {
		return false, errors.New("Invalid nick")
	}
	if len(contact.Mobile) == 0 {
		return false, errors.New("Each contact has to have at least one mobile phone number")
	}
	for _, mobile := range contact.Mobile {
		if !isMobile(mobile) {
			return false, errors.New("Invalid mobile phone number")
		}
	}
	if len(contact.Email) > 0 {
		for _, email := range contact.Email {
			if !isEmail(email) {
				return false, errors.New("Invalid email address")
			}
		}
	}
	return true, nil
}

//isEmail checks email address format
func isEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

//isMobile checks mobile phones format (spanish)
func isMobile(mobile string) bool {
	Re := regexp.MustCompile(`[6|7][0-9]{8}$`)
	return Re.MatchString(mobile)
}

//isValidNick verifies if a nick is from 3 to 36 alphanumeric characters long
func isValidNick(nick string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z0-9]{3,36}$`)
	return Re.MatchString(nick)
}
