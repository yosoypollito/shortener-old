package auth

import (
	"github.com/deta/deta-go/service/base"
	"shortener/api/database"
	"shortener/api/config"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"

	"time"
	"fmt"
)

type Dates struct {
	Creation time.Time `json:"creation"`
	Modified time.Time `json:"modified"`
	Lastlogin time.Time `json:"lastLogin"`
}

type User struct{
	Email string `json:"email"`
	Pass string `json:"pass"`
	Dates Dates `json:"dates"`
	Key string `json:"key"`
}

type TokenPayload struct{
	Exp time.Time `json:"exp"`
	User string `json:"user"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

type Form struct{
	Email string `json:"email" form:"email" binding:"required,email"`
	Pass string `json:"pass" form:"pass" binding:"required"`
}

type Register struct{
	Email string `json:"email" form:"email" binding:"required,email"`
	Pass string `json:"pass" form:"pass" binding:"required,eqfield=Cpass"`
	Cpass string `json:"cpass" form:"cpass" binding:"required,eqfield=Pass"`
}

func (self *Register) CheckPassMatch() bool {
	return self.Pass == self.Cpass
}

func CreateToken(payload *TokenPayload) (string, error){

	SECRET := []byte(config.EnvVariable("SECRET"))

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := newToken.SignedString(SECRET)

	if err != nil {
		fmt.Print(err,"\n")
		return "", err
	}

	return signedToken, nil;
}

func GetUserByEmail(email string) ([]*User, error){

	query := base.Query{
		{ "email": email },
	}

	results, err := database.GetFromBase[User]("Users", query)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func CreateUser(email string, pass string) (bool, error){

	userDb, err := database.Getbase("Users")

	if err != nil {
		return false, err
	}

	hashedPass, err := HashPass(pass)

	if err != nil {
		return false, nil
	}

	now := time.Now()

	newUser := User{
		Email:email,
		Pass:hashedPass,
		Dates: Dates{
			Creation: now,
			Modified: now,
			Lastlogin: now,
		},
	}

	_, err = userDb.Put(newUser)

	if err != nil {
		return false, err
	}

	return true, nil
}

func HashPass(pass string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	return string(bytes), err
}

func CheckPassHash(pass, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	return err == nil
}
