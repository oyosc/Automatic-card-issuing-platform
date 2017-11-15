package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"net/http"
	// "encoding/json"
	"automatic/backEnd/model"
)

func GetMd5String(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func  UserLogin(c *gin.Context) {
	var info model.UserLoginInfo
	var user model.User
	err := c.BindJSON(&info)
	fmt.Println("111111111")
	fmt.Println(err)
	if err != nil{
		fmt.Println(err)
	}else{
		name := info.Username
		passwd := info.Password
		if queryUserErr := model.DB.Where("name = ? and passwd = ?", name, passwd).Find(&user).Error; queryUserErr!=nil{
			c.JSON(200, gin.H{
				"status":10001,
				"message":"用户账号或者密码错误",
				})
		}else{
			session := sessions.Default(c)
			local_time := time.Now().UnixNano()/1e6
			md5_time := GetMd5String(string(int(local_time)))
			cookie_value := name + "|" + md5_time + "|" + passwd
	    session.Set("cookie", cookie_value)
	    session.Save()
			c.JSON(200, gin.H{
				"status":200,
				"user":user,
				})
		}
	}
}

func UserDetail(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
					"status_code": 200,
					"message": "获取用户详情",
	})
}

func GetUserByName(name string, passwd string) (int, error){
	type UserCountResult struct{
		UserCount int `json:"user_count"`
	}
	var userCountResult UserCountResult
	if err := model.DB.Raw("select count(*) as user_count from users where name = ? and passwd = ?", name, passwd).Scan(&userCountResult).Error; err == nil{
		fmt.Println("用户查询成功")
		if userCountResult.UserCount == 0{
			fmt.Println("没有该用户")
			return 0, nil
		}else{
			fmt.Println("该用户存在")
			return 1, nil
		}
	}else{
		fmt.Println("用户查询失败")
		return 0, err
	}

}