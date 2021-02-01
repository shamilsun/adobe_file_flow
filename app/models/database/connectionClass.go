package database

import (
	"../dotenv"
	"../log"
	//"../settings"
	"github.com/go-pg/pg"
	//	"github.com/go-pg/pg/orm"
)

func getDatabaseConnection() *pg.DB {

	options := pg.Options{
		User:     dotenv.GetEnv().DbUser,     //"postgres",          //postgres.User,
		Password: dotenv.GetEnv().DbPassword, //"123",               //postgres.Password,
		Database: dotenv.GetEnv().Database,   //"self_portrait_v1",  //postgres.Database,
		Addr:     dotenv.GetEnv().DbAddr,     //"10.16.66.214:5432", //postgres.Host,
		//Port:     "5432", //postgres.Port,
	}
	log.Println(options.User, options.Password, options.Addr)
	db := pg.Connect(&options)

	return db

}

func GetDB() *pg.DB {
	return getDatabaseConnection()
}

func CheckDB() {
	log.Println("Request connect to db")
	getDatabaseConnection()
	log.Println("Connected to DB")
}
