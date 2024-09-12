package main

import (
	"log"
	"net/http"

	//authdata "github.com/skuril-bobishku/test-task-backdev"

	"github.com/skuril-bobishku/test-task-backdev/config"
	"github.com/skuril-bobishku/test-task-backdev/internal/routing"
)

func main() {
	/*db, err := dbpostgres.ConnectDB(dbpostgres.DBCfg{
		Host:     config.Phost,
		Port:     config.Pport,
		Username: config.Pusername,
		DBName:   config.Pdbname,
		Password: config.Ppassword,
		SSLMode:  config.Psslmode,
	})
	if err != nil {
		log.Fatal(err)
	}

	user, err := dbpostgres.SearchUser(db, guid)

	fmt.Println("id: ", user.Guid,
	"username: " + user.Name,
	"email: " + user.Email,
	"ipadd: " + user.IPadd,
	"login: " + user.Login,
	"pass: " + user.Password,
	"session_id: ", user.Session_id)*/

	http.HandleFunc("/", routing.Test)
	http.HandleFunc("/create", routing.CreatePage)   // ?guid=
	http.HandleFunc("/refresh", routing.RefreshPage) //.Methods("GET")

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
