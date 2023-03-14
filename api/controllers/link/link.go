package link

import (
	"github.com/gin-gonic/gin"
	"shortener/api/structs/links"
	"net/http"
	"shortener/api/helpers"
	"shortener/api/config"
	"fmt"
)

func GetById(c *gin.Context){

	id := c.Param("id")

	links, err := links.Get(id)

	var scope string

	if err != nil {
		scope = config.EnvVariable("APP_URL") + "/app/"
		c.Redirect(http.StatusMovedPermanently, scope)
		return
	}

	scope = string(links.Scope)

	c.Redirect(http.StatusMovedPermanently, scope)
}

func CreateNew(c *gin.Context){

	var linkReq links.Body

	if err := helpers.BindJsonOrAbort(&linkReq, c); err != nil{
		return
	}

	link, err := links.Create(linkReq.Scope)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errors":helpers.GetErrors(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"link":link,
	})
}
