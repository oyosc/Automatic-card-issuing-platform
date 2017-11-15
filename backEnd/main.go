package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	// "net/http"
	"automatic/backEnd/config"
	"automatic/backEnd/model"
	"automatic/backEnd/controller/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"fmt"
)

func init(){
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	model.DB = db
}

func main() {
	r := gin.New()

	apiRouter := r.Group("/api/v1")
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(30*time.Minute),
		Path: "/",
	})
	apiRouter.Use(sessions.Sessions("mysession", store))
	apiRouter.Use(api.Auth())
	{
		apiRouter.GET("/user/detail", api.UserDetail)
		apiRouter.POST("/user/login", api.UserLogin)
	}
	r.Run(":8080")
}
