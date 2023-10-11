package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func main() {
	username := "zouying"
	token, err := genToken(username)
	if err != nil {
		log.Fatalf("generate token failed: %v", err)
	}

	log.Printf("username=%s token=%s", username, token)

	claims, err := parseToken(token)
	if err != nil {
		log.Fatalf("username=%s token=%s parse token error: %v", username, token, err)
	}
	log.Printf("username=%s token=%s claims: %+v", username, token, claims)

	// --- 等待延期 ---
	time.Sleep(TokenExpireDuration)
	time.Sleep(3 * time.Second)

	log.Printf("------------ After token expired ---------")
	claims, err = parseToken(token) // error: token is expired
	if err != nil {
		v, ok := err.(*jwt.ValidationError)
		if ok && v.Errors == jwt.ValidationErrorExpired {
			log.Printf("username=%s token=%s expired", username, token)
			return
		}

		log.Fatalf("username=%s token=%s parse token error: %v", username, token, err)
	}
	log.Printf("username=%s token=%s claims: %+v", username, token, claims)

}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// const TokenExpireDuration = time.Hour
const TokenExpireDuration = time.Second * 10

var MySecret = []byte("AllYourBase")

func genToken(username string) (string, error) {

	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(MySecret)
}

func parseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {

		// log.Fatalf("parse with clains error: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
