package server

/** HTTP router configuration **/

import (
	"encoding/json"
	"errors"
	"log"
	"mgo-n-goji/pkg"
	"net/http"
	"strings"

	"github.com/globalsign/mgo"
	goji "goji.io"
	"goji.io/pat"
)

//contactRouter type definition
type contactRouter struct {
	contactService root.ContactService
}

//NewContactRouter defines each path,handlerfunction pair
func NewContactRouter(cs root.ContactService, mux *goji.Mux) *goji.Mux {
	contactRouter := contactRouter{cs}
	mux.HandleFunc(pat.Get("/contacts"), contactRouter.getContacts)
	mux.HandleFunc(pat.Get("/contacts/:nick"), contactRouter.getContact)
	mux.HandleFunc(pat.Post("/contacts"), contactRouter.addContact)
	mux.HandleFunc(pat.Put("/contacts/:nick"), contactRouter.updateContact)
	mux.HandleFunc(pat.Delete("/contacts/:nick"), contactRouter.deleteContact)
	mux.HandleFunc(pat.Patch("/contacts/:nick/addMobile/:mobile"), contactRouter.addMobile)
	mux.HandleFunc(pat.Patch("/contacts/:nick/delMobile/:mobile"), contactRouter.delMobile)
	mux.HandleFunc(pat.Patch("/contacts/:nick/addEmail/:email"), contactRouter.addEmail)
	mux.HandleFunc(pat.Patch("/contacts/:nick/delEmail/:email"), contactRouter.delEmail)
	mux.HandleFunc(pat.Patch("/contacts/:oldNick/newNick/:newNick"), contactRouter.changeNick)
	return mux
}

//getContact handler returns a contact according to its nick
func (cr *contactRouter) getContact(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	contact, err := cr.contactService.GetContact(nick)
	if err != nil {
		switch err {
		default:
			log.Println("getContact handler has failed: ", err.Error())
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNoContent)
			return
		}
	}
	JSON(w, http.StatusOK, contact)
}

//addContact handler adds a contact
func (cr *contactRouter) addContact(w http.ResponseWriter, r *http.Request) {
	contact, err := decodeContact(r)
	if err != nil {
		log.Println("Invalid request payload", err.Error())
		Code(w, http.StatusBadRequest)
		return
	}
	err = cr.contactService.AddContact(&contact)
	if err != nil {
		log.Println(err.Error())
		Code(w, http.StatusBadRequest)
		return
	}
	Code(w, http.StatusCreated)
}

//getContacts handler gets all contacts
func (cr *contactRouter) getContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := cr.contactService.GetContacts()
	if err != nil {
		log.Println("getContacts handler has failed: ", err.Error())
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(contacts) == 0 {
		Code(w, http.StatusNoContent)
		return
	}
	JSON(w, http.StatusOK, contacts)
}

//updateContact handler updates a contact identified by nick
func (cr *contactRouter) updateContact(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	contact, err := decodeContact(r)
	if err != nil {
		log.Println("Invalid request payload", err.Error())
		Code(w, http.StatusBadRequest)
		return
	}
	err = cr.contactService.UpdateContact(nick, &contact)
	if err != nil {
		switch err {
		default:
			log.Println("updateContact handler has failed: ", err.Error())
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//deleteContact handler deletes a contact identified by nick
func (cr *contactRouter) deleteContact(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	err := cr.contactService.DeleteContact(nick)
	if err != nil {
		switch err {
		default:
			log.Println("deleteContact handler has failed: ", err.Error())
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//addMobile handler adds a mobile to the contact identified by nick
func (cr *contactRouter) addMobile(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	mobile := pat.Param(r, "mobile")
	err := cr.contactService.AddMobile(nick, mobile)
	if err != nil {
		switch err {
		default:
			log.Println(err.Error())
			if strings.Index(err.Error(), "Invalid") == 0 {
				Code(w, http.StatusBadRequest)
				return
			}
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//delMobile handler removes a mobile to the contact identified by nick
func (cr *contactRouter) delMobile(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	mobile := pat.Param(r, "mobile")
	err := cr.contactService.DelMobile(nick, mobile)
	if err != nil {
		switch err {
		default:
			log.Println(err.Error())
			if strings.Index(err.Error(), "Invalid") == 0 {
				Code(w, http.StatusBadRequest)
				return
			}
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//addEmail handler adds an email address to the contact identified by nick
func (cr *contactRouter) addEmail(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	email := pat.Param(r, "email")
	err := cr.contactService.AddEmail(nick, email)
	if err != nil {
		switch err {
		default:
			log.Println(err.Error())
			if strings.Index(err.Error(), "Invalid") == 0 {
				Code(w, http.StatusBadRequest)
				return
			}
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//delEmail handler removes an email address to the contact identified by nick
func (cr *contactRouter) delEmail(w http.ResponseWriter, r *http.Request) {
	nick := pat.Param(r, "nick")
	email := pat.Param(r, "email")
	err := cr.contactService.DelEmail(nick, email)
	if err != nil {
		switch err {
		default:
			log.Println(err.Error())
			if strings.Index(err.Error(), "Invalid") == 0 {
				Code(w, http.StatusBadRequest)
				return
			}
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + nick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//changeNick handler sets a newNick the contact identified by oldNick
func (cr *contactRouter) changeNick(w http.ResponseWriter, r *http.Request) {
	oldNick := pat.Param(r, "oldNick")
	newNick := pat.Param(r, "newNick")
	err := cr.contactService.ChangeNick(oldNick, newNick)
	if err != nil {
		switch err {
		default:
			log.Println(err.Error())
			if strings.Index(err.Error(), "Invalid") == 0 {
				Code(w, http.StatusBadRequest)
				return
			}
			Error(w, http.StatusInternalServerError, err.Error())
			return
		case mgo.ErrNotFound:
			log.Println("Contact " + oldNick + " not found")
			Code(w, http.StatusNotFound)
			return
		}
	}
	Code(w, http.StatusOK)
}

//decodeContact function transforms a contact to its json representation
func decodeContact(r *http.Request) (root.Contact, error) {
	var c root.Contact
	if r.Body == nil {
		return c, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	return c, err
}
