package main

import (
	// "github.com/itcupnvyk/backend-class/database/model"
	"log"
	"database/sql"
	"github.com/itcupnvyk/backend-class/database/repository"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/user_service")
	db.Exec(`CREATE TABLE IF NOT EXISTS users(id bigint primary key, name varchar(255))`)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)

	// user := &model.User{
	// 	ID: 2,
	// 	Name: "User Two",
	// }
	// err = userRepo.Create(user)
	// if err != nil {
	// 	log.Println(err)
	// }

	// users, err := userRepo.FindAll()
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(users)

	user, err := userRepo.FindByID(1)
	if err != nil {
		log.Println(err)
	}
	log.Println(user)
}