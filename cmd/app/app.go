package main

/** App **/

import (
	"log"
	"mgo-n-goji/pkg"
	"mgo-n-goji/pkg/config"
	"mgo-n-goji/pkg/mongo"
	"mgo-n-goji/pkg/server"
)

//App type definition
type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

//Initialize reads the configuration and open a new session to mongoDB.
//It fails, logs it.
//It succeeds, starts a contact service instance, then configure the http server.
func (a *App) Initialize() {
	a.config = config.GetConfig()
	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	cs := mongo.NewContactService(a.session.Copy(), a.config.Mongo)
	a.server = server.NewServer(cs, a.config)
}

//Run start the http server
func (a *App) Run() {
	defer a.session.Close()
	a.server.Start()
}
