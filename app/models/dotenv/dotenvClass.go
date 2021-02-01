package dotenv

import (
	"log"
	"os"
)

import "github.com/joho/godotenv"

var env IDotEnv
var isLoaded = false

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
		panic(err)
	}

	env = IDotEnv{
		Database:   os.Getenv("db.name"),
		DbAddr:     os.Getenv("db.addr"),
		DbUser:     os.Getenv("db.user"),
		DbPassword: os.Getenv("db.password"),

		ClientId:     os.Getenv("sendpulse.clientId"),
		ClientSecret: os.Getenv("sendpulse.clientSecret"),
		FromName:     os.Getenv("sendpulse.fromName"),
		FromEmail:    os.Getenv("sendpulse.fromEmail"),

		QuickEmailVerification: os.Getenv("quickemailverification.api"),

		Mode: os.Getenv("mode"),
	}

	isLoaded = true
}

func GetEnv() *IDotEnv {
	if !isLoaded {
		InitEnv()
	}
	return &env
}
