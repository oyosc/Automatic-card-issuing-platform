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
	"errors"
)


func Auth() gin.HandlerFunc{
	return func(c *gin.Context){
		var status_code int
		var status_err error
		fmt.Println(c.Request.URL)
		url := c.Request.URL.Path
		if needCheckAuth(url){
			session := sessions.Default(c)
			userAccessToken := session.Get("cookie")
			fmt.Println(userAccessToken)
			if userAccessToken == nil{
				status_code = 403
				status_err = errors.New("cookie not found or wrong")
			}else{
				splitToken := strings.Split(userAccessToken.(string), "|")
				userName := splitToken[0]
				passwd := splitToken[2]
				count, err := GetUserByName(userName, passwd)
				if(err != nil){
					status_code = 500
					status_err = err
				}else{
					status_code = 200
					if count == 1{
						status_err = nil
					}else{
						status_err = errors.New("err_code:10001,该用户不存在") //用户不存在
					}
					
				}
			}
			if status_code != 200{
				c.JSON(http.StatusOK, gin.H{
					"status_code": status_code,
					"error": status_err.Error(),
				})
				c.Abort()
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