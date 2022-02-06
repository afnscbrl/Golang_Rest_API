package controller

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/afnscbrl/Golang_Rest_API/database"
	"github.com/afnscbrl/Golang_Rest_API/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key") //GET ENV "string"

type Users_Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SHA256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users
	var users []models.Users
	json.NewDecoder(r.Body).Decode(&newUser)
	database.DB.Find(&users)

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].Username)
		if newUser.Username == users[i].Username {
			log.Println("Error 409: Conflict - This username already exist")
			http.Error(w, "Error 409: Conflict - This username already exist", http.StatusConflict)
			return
		}
	}
	newUser.Passwordhash = SHA256Encoder(newUser.Passwordhash)
	database.DB.Create(&newUser)
	json.NewEncoder(w).Encode(newUser.Username)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestCredentials models.Users
	var credentials []models.Users
	var user_auth = make(map[string]string)

	json.NewDecoder(r.Body).Decode(&requestCredentials)
	database.DB.Find(&credentials)

	for i := 0; i < len(credentials); i++ {
		if requestCredentials.Username == credentials[i].Username {

			if credentials[i].Passwordhash != SHA256Encoder(requestCredentials.Passwordhash) {
				http.SetCookie(w,
					&http.Cookie{
						Name:    "token",
						Value:   "",
						Expires: time.Now().Add(time.Second * 2),
					})
				break
			}
			user_auth[credentials[i].Username] = requestCredentials.Passwordhash
			expirationTime := time.Now().Add(time.Minute * 5)

			claims := &Users_Claims{
				Username: user_auth[requestCredentials.Username],
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w,
				&http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
				})
			break
		}
	}

	w.WriteHeader(http.StatusUnauthorized)

}
