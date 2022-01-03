package controllers

import (
	"encoding/json"
	"golang-rest-api/common"
	"golang-rest-api/models"
	"golang-rest-api/secure"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TokenI struct {
	Token string `json:"token"`
}

func Login(writer http.ResponseWriter, request *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := models.Users{}

	// Json Body User
	creds := &Credentials{}
	error := json.NewDecoder(request.Body).Decode(&creds)
	if error != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	db := common.GetConnection()
	defer db.Close()

	// Json Db User
	db.Where("email = ?", creds.Email).First(&user)

	matchPassword := secure.CheckPasswordHash(creds.Password, user.Password)

	if !matchPassword {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err_token := secure.CreateToken(uint64(user.ID))

	if err_token != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	tokenData := TokenI{Token: token}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tokenData)
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
