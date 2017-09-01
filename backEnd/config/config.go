package config

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)


var jsonData map[string] interface{}

func initJSON(){
	bytes, err := ioutil.ReadFile("../config.json")
	if err != nil{
		fmt.Println("ReadFile:", err.Error())
	}
	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr = reg.ReplaceAllString(configStr, "")
	fmt.Println(reg)
	bytes = []byte(configStr)
	if err := json.Unmarshal(bytes, &jsonData); err !=nil{
		fmt.Println("invalid config:", err.Error())
	}
}

type dbConfig struct{
	Dialect string
	Database string
	User string
	Password string
	Host string
	Port int
	Charset string
	URL string
	MaxIdleConns int
	MaxOpenConns int
}

var DBConfig dbConfig

func initDB(){
	utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	portStr := fmt.Sprintf("%d", DBConfig.Port)
	url := "host={host} user={user} dbname={dbname} sslmode=disable password={password}"
	url = strings.Replace(url, "{dbname}", DBConfig.Database, -1)
	url = strings.Replace(url, "{user}", DBConfig.User, -1)
	url = strings.Replace(url, "{host}", DBConfig.Host, -1)
	url = strings.Replace(url, "{password}", DBConfig.Password, -1)
	DBConfig.URL = url
	fmt.Println(DBConfig.URL)
}

func init(){
	initJSON()
	// fmt.Println(jsonData)
}
























