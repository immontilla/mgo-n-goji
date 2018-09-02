package config

/** Configuration **/

import (
	"mgo-n-goji/pkg"
	"os"
)

//GetConfig returns the app configuration settings
func GetConfig() *root.Config {
	return &root.Config{
		Mongo: &root.MongoConfig{
			IP:     envOrDefault("MGONGOJI_MONGODB", "127.0.0.1:27017"),
			DbName: envOrDefault("MGONGOJI_DBNAME", "phonebook")},
		Server: &root.ServerConfig{Port: envOrDefault("MGONGOJI_IPPORT", ":9889")}}
}

//envOrDefault gets an environment variable value, returning a default value if is not set
func envOrDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}
	return value
}
