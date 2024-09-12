package routing

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/skuril-bobishku/test-task-backdev/config"
	"github.com/skuril-bobishku/test-task-backdev/internal/auth"
	"github.com/skuril-bobishku/test-task-backdev/internal/dbpostgres"
)

func Test(w http.ResponseWriter, r *http.Request) {}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	guid, err := strconv.Atoi(r.URL.Query().Get("guid"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbpostgres.ConnectDB(dbpostgres.DBCfg{
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

	// ПОИСК ПОЛЬЗОВАТЕЛЯ
	user, err := dbpostgres.SearchUser(db, guid)
	if err != nil {
		log.Fatal(err)
	}

	// как брать IPv4 не знаю
	//"CF-Connecting-IP", "X-Forwarded-For" и "X-Real-IP" не брал, потому что их подменять можно https://stackoverflow.com/a/55790/1584308
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	createTime := time.Now()

	accessToken, refreshToken := auth.GeneratePair(user.Name, createTime, ip)
	bcryptRefreshToken, base64RefreshToken := auth.CryptToken(refreshToken)

	err = dbpostgres.InsertIPAddress(db, guid, ip)
	if err != nil {
		log.Fatal(err)
	}

	session_id, err := dbpostgres.InsertRefresh(db, guid, bcryptRefreshToken, createTime)
	if err != nil {
		log.Fatal(err)
	}

	ac := &http.Cookie{
		Name:     "access-cookie",
		Value:    accessToken,
		Expires:  createTime.Add(config.ExpAccess),
		HttpOnly: true,
	}

	refresh_value := base64RefreshToken + "?" + strconv.Itoa(session_id)

	rc := &http.Cookie{
		Name:     "refresh-cookie",
		Value:    refresh_value,
		Expires:  createTime.Add(config.ExpAccess),
		HttpOnly: true,
	}

	http.SetCookie(w, ac)
	http.SetCookie(w, rc)
}

func RefreshPage(w http.ResponseWriter, r *http.Request) {
	ref_string := r.URL.Query().Get("token")
	parts := strings.Split(ref_string, "?")

	session_id, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	byteRefreshToken, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	userRefreshToken := string(byteRefreshToken)

	db, err := dbpostgres.ConnectDB(dbpostgres.DBCfg{
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

	dbRefreshToken, err := dbpostgres.SearchSession(db, session_id)
	if err != nil {
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbRefreshToken), []byte(userRefreshToken))
	if err != nil {
		fmt.Println("Not valid refresh-token")
	}

	ip_bd, err := dbpostgres.SearchIPadd(db, session_id)
	if err != nil {
		log.Fatal(err)
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	if ip_bd != ip {
		fmt.Println("send mail") // не разбирался с отправкой почты
	}

	fmt.Println(userRefreshToken)
}
