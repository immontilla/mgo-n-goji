package mongo

/** Contact service **/

import (
	"errors"
	"fmt"
	"mgo-n-goji/pkg"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//ContactService type definition
type ContactService struct {
	collection *mgo.Collection
}

//NewContactService creates a contact service instance
func NewContactService(session *mgo.Session, config *root.MongoConfig) *ContactService {
	fmt.Println("Accessing " + config.DbName + " database")
	collection := session.DB(config.DbName).C("contacts")
	collection.EnsureIndex(contactIndexByNick())
	collection.EnsureIndex(contactIndexByMobile())
	return &ContactService{collection}
}

//GetContact gets a contact according to its nick from the collection
func (cs *ContactService) GetContact(nick string) (root.Contact, error) {
	contact := root.Contact{}
	err := cs.collection.Find(bson.M{"nick": nick}).One(&contact)
	return contact, err
}

//AddContact adds a contact to the collection
func (cs *ContactService) AddContact(contact *root.Contact) error {
	isValid, validationError := isValidContact(contact)
	if !isValid {
		return validationError
	}
	err := cs.collection.Insert(&contact)
	return err
}

//GetContacts returns all contacts from the collection
func (cs *ContactService) GetContacts() ([]root.Contact, error) {
	var contacts []root.Contact
	err := cs.collection.Find(nil).All(&contacts)
	return contacts, err
}

//UpdateContact updates a contact
func (cs *ContactService) UpdateContact(nick string, contact *root.Contact) error {
	err := cs.collection.Update(bson.M{"nick": nick}, &contact)
	return err
}

//DeleteContact removes a contact from the collection
func (cs *ContactService) DeleteContact(nick string) error {
	err := cs.collection.Remove(bson.M{"nick": nick})
	return err
}

//AddMobile adds a mobile to the contact identified by nick
func (cs *ContactService) AddMobile(nick string, mobile string) error {
	if !isMobile(mobile) {
		return errors.New("Invalid mobile")
	}
	match := bson.M{"nick": nick}
	contact := root.Contact{}
	errFindOne := cs.collection.Find(match).One(&contact)
	if errFindOne == nil {
		for _, mob := range contact.Mobile {
			if mob == mobile {
				return errors.New("Invalid mobile phone number. Already exists")
			}
		}
	}
	change := bson.M{"$push": bson.M{"mobile": mobile}}
	err := cs.collection.Update(match, change)
	return err
}

//DelMobile removes a mobile from the contact identified by nick
func (cs *ContactService) DelMobile(nick string, mobile string) error {
	if !isMobile(mobile) {
		return errors.New("Invalid mobile")
	}
	match := bson.M{"nick": nick}
	contact := root.Contact{}
	errFindOne := cs.collection.Find(match).One(&contact)
	if errFindOne == nil {
		if len(contact.Mobile) == 1 && contact.Mobile[0] == mobile {
			return errors.New("Invalid mobile, a contact has to have at least one mobile")
		}
	}
	change := bson.M{"$pull": bson.M{"mobile": mobile}}
	err := cs.collection.Update(match, change)
	return err
}

//AddEmail adds an email address to the contact identified by nick
func (cs *ContactService) AddEmail(nick string, email string) error {
	if !isEmail(email) {
		return errors.New("Invalid email address")
	}
	match := bson.M{"nick": nick}
	contact := root.Contact{}
	errFindOne := cs.collection.Find(match).One(&contact)
	if errFindOne == nil {
		for _, emailAddress := range contact.Email {
			if emailAddress == email {
				return errors.New("Invalid email address. Already exists")
			}
		}
	}
	change := bson.M{"$push": bson.M{"email": email}}
	err := cs.collection.Update(match, change)
	return err
}

//DelEmail removes an email address from the contact identified by nick
func (cs *ContactService) DelEmail(nick string, email string) error {
	if !isEmail(email) {
		return errors.New("Invalid email address")
	}
	match := bson.M{"nick": nick}
	change := bson.M{"$pull": bson.M{"email": email}}
	err := cs.collection.Update(match, change)
	return err
}

//ChangeNick changes the nick of a contact identified by oldNick
func (cs *ContactService) ChangeNick(oldNick string, newNick string) error {
	if !isValidNick(newNick) {
		return errors.New("Invalid new nick")
	}
	match := bson.M{"nick": oldNick}
	change := bson.M{"$set": bson.M{"nick": newNick}}
	err := cs.collection.Update(match, change)
	return err
}
