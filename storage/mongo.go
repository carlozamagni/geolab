package storage


import (
	"errors"
	"gopkg.in/mgo.v2"
	s "strings"
)


type dbConfig struct{
	env string
	url string
}

func loadConfiguration(env string) (config dbConfig, err error) {

	switch s.ToLower(env) {

	case "local":
		return dbConfig{
			env:"local",
			url:"mongodb://127.0.0.1:27017/admin"},
		nil

	case "dev":
		return dbConfig{
			env:"dev",
			url:""},
		nil

	case "test":
		return dbConfig{
			env:"test",
			url:""},
		nil

	case "prod":
		return dbConfig{
			env:"prod",
			url:""},
		nil
	}

	return dbConfig{env:"", url:""}, errors.New("configuration not found")
}

func MongoSession(env string) *mgo.Session{
	config, err := loadConfiguration(env)
	s, err := mgo.Dial(config.url)

	if err != nil {
		panic(err)
	}

	return s
}