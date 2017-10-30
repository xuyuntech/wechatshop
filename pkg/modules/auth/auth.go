package auth

import (
	"github.com/dgrijalva/jwt-go"
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

type (
	Authenticator interface{
		Authenticate(username, password, hash string) (bool, error)
		GenerateToken(tokenKey, uid string, username string, exp ...int64) (string, error)
		IsUpdateSupported() bool
		Name() string
	}
)

func Hash(data string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(h[:]), err
}

func HashPassword(password, salt string) (string, error) {
	h, err := scrypt.Key([]byte(password),[]byte(salt), 16384, 8, 1, 3)
	return string(h[:]), err
}

func GenerateToken(tokenKey, uid string, username string, exp ...int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["uid"] = uid
	claims["username"] = username
	if len(exp) > 0 {
		claims["exp"] = exp[0]
	} else {
		claims["exp"] = time.Now().Add(time.Hour * 480).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenKey))
}

func AuthToken(tokenKey, tokenString string) (bool, map[string]interface{}, error) {
	if tokenString == "" {
		return false, nil, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return []byte(tokenKey), nil
	})
	if err != nil {
		return false, nil ,err
	}
	if !token.Valid {
		return false, nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return false, nil, errors.New("token parse err")
	}

	return true, map[string]interface{}{
		"uid": claims["uid"],
		"username": claims["username"],
	}, nil

}