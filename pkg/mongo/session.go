package mongo

/** MongoDB Session **/

import (
	"fmt"
	"mgo-n-goji/pkg"

	"github.com/globalsign/mgo"
)

//Session type definition
type Session struct {
	session *mgo.Session
}

//NewSession creates a new session to mongoDB
func NewSession(config *root.MongoConfig) (*Session, error) {
	fmt.Println("Opening session to mongoDB at " + config.IP)
	session, err := mgo.Dial(config.IP)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &Session{session}, err
}

//Copy returns a mongoDB session
func (s *Session) Copy() *mgo.Session {
	return s.session.Copy()
}

//Close closes a mongoDB session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

//DropDatabase drops a mongoDB database named {db}
func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
