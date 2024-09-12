package dbpostgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	authdata "github.com/skuril-bobishku/test-task-backdev"
)

type DBCfg struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectDB(dbcfg DBCfg) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbcfg.Host, dbcfg.Port, dbcfg.Username, dbcfg.DBName, dbcfg.Password, dbcfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertUser(db *sqlx.DB, username string, email string, login string, password string) (int, error) {
	var gu_id int

	err := db.QueryRow("INSERT INTO users (username, email, login, pass) VALUES ($1, $2, $3, $4) RETURNING gu_id;",
		username, email, login, password).Scan(&gu_id)

	if err != nil {
		return -1, err
	}

	return gu_id, nil
}

func SearchUser(db *sqlx.DB, gu_id int) (authdata.AuthData, error) {
	var user authdata.AuthData

	user.Guid = gu_id

	err := db.QueryRow("SELECT username, email, login, pass FROM users WHERE gu_id = $1;",
		gu_id).Scan(&user.Name, &user.Email, &user.Login, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertIPAddress(db *sqlx.DB, gu_id int, ipadd string) error {
	_, err := db.Exec("UPDATE users SET ipadd = $2 WHERE gu_id = $1;", gu_id, ipadd)

	if err != nil {
		return err
	}

	return nil
}

func InsertRefresh(db *sqlx.DB, gu_id int, refreshToken string, exp time.Time) (int, error) {
	var s_id int

	err := db.QueryRow("INSERT INTO sessions (refresh_crypt, exp_time) VALUES ($1, $2) RETURNING s_id;",
		refreshToken, exp).Scan(&s_id)

	if err != nil {
		return -1, err
	}

	_, err = db.Exec("UPDATE users SET session_id = $2 WHERE gu_id = $1;", gu_id, s_id)

	if err != nil {
		return -1, err
	}

	return s_id, nil
}

func SearchSession(db *sqlx.DB, session_id int) (string, error) {
	var refresh_crypt string

	err := db.QueryRow("SELECT refresh_crypt FROM sessions WHERE s_id = $1;", session_id).Scan(&refresh_crypt)
	if err != nil {
		return "", err
	}

	return refresh_crypt, nil
}

func SearchRefresh(db *sqlx.DB, refresh string) (int, time.Time, error) {
	var s_id int
	var exp time.Time

	err := db.QueryRow("SELECT s_id, exp_time FROM sessions WHERE refresh_crypt = $1;", refresh).Scan(&s_id, &exp)
	if err != nil {
		return -1, time.Time{}, err
	}

	return s_id, exp, nil
}

func SearchIPadd(db *sqlx.DB, session_id int) (string, error) {
	var ipadd string

	err := db.QueryRow("SELECT ipadd FROM users WHERE session_id = $1;", session_id).Scan(&ipadd)
	if err != nil {
		return "", err
	}

	return ipadd, nil
}
