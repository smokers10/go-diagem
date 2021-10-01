package jwt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	ID    int
	Email string
	Type  string
}

type PayloadResetPassword struct {
	Email     string
	Type      string
	Code      string
	ExpiredAt int
}

func signKey() []byte {
	if os.Getenv("PRODUCTION_MODE") == "local" || os.Getenv("PRODUCTION_MODE") == "development" || os.Getenv("PRODUCTION_MODE") == "" {
		return []byte("this_is_local_secret")
	}
	return []byte(os.Getenv("SECRET"))
}

func Sign(payload *Payload) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = payload.ID
	claims["email"] = payload.Email
	claims["type"] = payload.Type
	claims["exp"] = time.Now().Add((time.Hour * 24) * 7)

	tokenString, err := token.SignedString(signKey())
	if err != nil {
		panic(err)
	}

	return tokenString
}

func Verify(tokenString string) (payload *Payload) {
	splitedToken := strings.Split(tokenString, "Bearer ")
	token, err := jwt.Parse(splitedToken[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ada yang salah dengan token otentifikasi")
		}

		return signKey(), nil
	})
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	claims := token.Claims.(jwt.MapClaims)
	payload.ID, _ = strconv.Atoi(claims["id"].(string))
	payload.Email = claims["email"].(string)
	payload.Type = claims["type"].(string)

	return payload
}

func SignResetPassword(payload *PayloadResetPassword) string {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = payload.Email
	claims["type"] = payload.Type
	claims["code"] = payload.Code
	claims["expired_at"] = payload.ExpiredAt

	tokenString, err := token.SignedString(signKey())
	if err != nil {
		panic(err)
	}

	return tokenString
}

func VerifyResetPessword(tokenString string) (payload *PayloadResetPassword) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ada yang salah dengan token reset password, silahkan buat lagi")
		}

		return signKey(), nil
	})
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	claims := token.Claims.(jwt.MapClaims)
	payload.Email = claims["email"].(string)
	payload.Type = claims["type"].(string)
	payload.Code = claims["code"].(string)
	payload.ExpiredAt, _ = strconv.Atoi(claims["expired_at"].(string))

	return payload
}
