package mongo

/** Contact collection schema **/

import (
	"github.com/globalsign/mgo"
)

//contactIndexByNick defines an index to the Nick field
func contactIndexByNick() mgo.Index {
	return mgo.Index{
		Key:        []string{"nick"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

//contactIndexByMobile defines an index to the Mobile field
func contactIndexByMobile() mgo.Index {
	return mgo.Index{
		Key:        []string{"mobile"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}
