package auth 

import (
	"shortener/api/helpers"
	"shortener/api/config"
	"shortener/api/structs/auth"

	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"fmt"
)

func Login(c *gin.Context){

	var userForm auth.Form

	if err := helpers.BindJsonOrAbort(&userForm, c); err != nil {
		return
	}

	user, err := auth.GetUserByEmail(userForm.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors":nil,
			"message":err.Error(),
		})
		return
	}

	if len(user) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors":nil,
			"message":"User not found",
		})
		return
	}
	
	passMatch := auth.CheckPassHash(userForm.Pass, user[0].Pass);

	if !passMatch {
		errorFields := []helpers.ErrorMsg { 
			{ Field: "Pass", Message: "Do not match",},
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors":errorFields,
		})

		return
	}

	fmt.Print(helpers.StringFromJson(userForm))

	tokenPayload := auth.TokenPayload{
		Exp: time.Now(),
		User: user[0].Key,
		Type: "authorization",
	}

	tokenString, err := auth.CreateToken(&tokenPayload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errors":fmt.Sprintf("Server Internal Error: %s",err.Error()),
		})
		return
	}

	c.SetCookie("Authorization", fmt.Sprintf("Bearer %v", tokenString), 0, "/", config.EnvVariable("APP_URL"), true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Authorized"})
}

func Register(c *gin.Context){
	var registerForm auth.Register

	if err := helpers.BindJsonOrAbort(&registerForm, c); err != nil {
		return
	}

	users, err := auth.GetUserByEmail(registerForm.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors":nil,
			"message":err.Error(),
		})
		return
	}

	if len(users) > 0 {

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors":nil,
			"message":"User already registered",
		})

		return
	}

	//We dont handle pass equal cpass cause binding already do that
	created, err := auth.CreateUser(registerForm.Email, registerForm.Pass)

	if !created {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors":nil,
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"errors":nil,
		"message":"User created",
	})
}

func Auth(c *gin.Context){
	auth, err := c.Cookie("Authorization")

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"auth":auth,
	})
}

func Logout(c *gin.Context){

	c.SetCookie("Authorization", "", -1, "/", "", true, true)

	cok, err := c.Cookie("Authorization")

	if err != nil{
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"cookie": cok,
	})

	scope := config.EnvVariable("APP_URL") + "/app/"
	c.Redirect(http.StatusOK, scope)
}
