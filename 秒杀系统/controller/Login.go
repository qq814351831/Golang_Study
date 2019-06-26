package controller

import (
	"Golang_Study/秒杀系统/database"
	"Golang_Study/秒杀系统/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)


func Login(c *gin.Context)(){
	id := 1
	var user model.User//结构体绑定
	err := database.DBsql.First(&user, id)//根据主键查询
	if err.Error != nil{
		log.Fatal(err.Error)
		os.Exit(1)
	}
	fmt.Println(user)
}
