package config

import (
	"time"

	authdata "github.com/skuril-bobishku/test-task-backdev"
)

var Port string = "8080"

var AccessKey = []byte("Acc secret")
var RefreshKey = []byte("Ref secret")

var ExpAccess time.Duration = time.Minute * 15
var ExpRefresh time.Duration = time.Hour * 72

// var BDconnect = "user=postgres dbname=testtaskbd sslmode=disable password=testtaskbd"
// надо подгружать из файла в .gitignore (godotenv)
var Phost = "localhost"
var Pport = "8090"
var Pusername = "postgres"
var Ppassword = "testtaskbd"
var Pdbname = "postgres"
var Psslmode = "disable"

var User = authdata.AuthData{
	Guid:     0,
	Name:     "Alex",
	Email:    "test@mail.com",
	IPadd:    "",
	Login:    "alex",
	Password: "pass",
}
