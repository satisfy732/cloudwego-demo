package main

import (
	"log"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/joho/godotenv"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
)
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dal.Init()
	mysql.DB.Create(&model.Student{Name:"test", ID:"test", Age:1, Class:"test"})
}