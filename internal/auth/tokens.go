package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/skuril-bobishku/test-task-backdev/config"
)

func GenerateAccessToken(name string, exp time.Time, ip string, refresh string) string {
	aToken := jwt.New(jwt.SigningMethodHS512)

	claims := aToken.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["exp"] = exp.Add(config.ExpAccess).Unix() // TODO: config.go
	claims["ip"] = ip
	claims["refresh"] = refresh

	signToken, err := aToken.SignedString(config.AccessKey)

	if err != nil {
		log.Fatal(err)
	}

	return signToken
}

func GenerateRefreshToken(exp time.Time, ip string) string {
	/*rToken := jwt.New(jwt.SigningMethodNone) // брал jwt, т.к. нужен ip, а так можно "случ. строка + . + хэш строки ip"

	claims := rToken.Claims.(jwt.MapClaims)
	claims["exp"] = exp.Unix()
	claims["ip"] = ip

	signToken, err := rToken.SignedString(jwt.UnsafeAllowNoneSignatureType)

	if err != nil {
		log.Fatal(err)
	}

	return signToken // <72 символов*/

	payload := fmt.Sprintf(`"exp": "%d","ip": "%s"`, exp.Unix(), ip)
	rToken := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(rToken[:])
}

func GeneratePair(name string, exp time.Time, ip string) (string, string) {
	refresh := GenerateRefreshToken(exp, ip)
	return GenerateAccessToken(name, exp, ip, refresh), refresh
}

func CryptToken(rt string) (string, string) {
	btb, err := bcrypt.GenerateFromPassword([]byte(rt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	bcryptRefreshToken := string(btb)

	return bcryptRefreshToken, base64.StdEncoding.EncodeToString([]byte(rt))
}

/*
	str := base64.StdEncoding.EncodeToString([]byte("Hello, playground"))
    fmt.Println(str)

    data, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
            log.Fatal("error:", err)
    }

    fmt.Printf("%q\n", data)
*/
