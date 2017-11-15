package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	// "net/http"
	// "encoding/json"
	// "io/ioutil"
	"fmt"
	"strings"
	// "automatic/backEnd/model"
	"net/http"
)


func Auth() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println(c.Request.URL)
		url := c.Request.URL.Path
		if needCheckAuth(url){
			session := sessions.Default(c)
			userAccessToken := session.Get("cookie")
			splitToken := strings.Split(userAccessToken.(string), "|")
			userName := splitToken[0]
			passwd := splitToken[2]
			status, err := GetUserByName(userName, passwd)
			fmt.Println(status)
			if(err != nil){
				c.JSON(http.StatusOK, gin.H{
					"status_code": 500,
					"error": err,
					})
			}else{
				c.Next()
			}

		}else{
			fmt.Println("there is not auth")
		}
	}
}

func needCheckAuth(url string) bool{
	switch{
		case url == "/api/v1/user/login":
			return false
		default:
			return true
	}
}