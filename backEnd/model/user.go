package model

type User struct{
	Id int `gorm:"primary_key" json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Password string `json:"name" binding:"required"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Status int `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}


