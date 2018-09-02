package root

/** Main types definitions **/

//Contact type definition
type Contact struct {
	Nick   string   `json:"nick"`
	Mobile []string `json:"mobile"`
	Email  []string `json:"email,omitempty"`
}

//ContactService interface definition
type ContactService interface {
	GetContacts() ([]Contact, error)
	GetContact(nick string) (Contact, error)
	AddContact(contact *Contact) error
	UpdateContact(nick string, contact *Contact) error
	DeleteContact(nick string) error
	AddMobile(nick string, mobile string) error
	DelMobile(nick string, mobile string) error
	AddEmail(nick string, email string) error
	DelEmail(nick string, email string) error
	ChangeNick(oldNick string, newNick string) error
}
