package root

/** Configuration type definitions **/

//MongoConfig tyoe definition (MongoDB settings)
type MongoConfig struct {
	IP     string
	DbName string
}

//ServerConfig type definitions (Server settings)
type ServerConfig struct {
	Port string
}

//Config type definition (App settings)
type Config struct {
	Mongo  *MongoConfig  `json:"mongo"`
	Server *ServerConfig `json:"server"`
}
