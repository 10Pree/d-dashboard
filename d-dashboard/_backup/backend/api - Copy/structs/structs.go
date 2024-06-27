package structs

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fname  string `json:"fname"`
	Lname  string `json:"lname"`
	Avatar string `json:"avatar"`
	Email  string `json:"emali"`
	OnisDelete int `json:"on_is_delete"`
}

type ReqCreateUser struct{
	Fname  string `json:"fname"`
	Lname  string `json:"lname"`
	Avatar string `json:"avatar"`
	Email  string `json:"emali"`
}

type ResCreateUser struct{
	Message string `json:"Message"`
	Status string `json:"Status"`
}