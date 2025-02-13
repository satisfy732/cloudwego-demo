package mysql

import (
	"fmt"
	"os"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/cloudwego/biz-demo/gomall/app/user/model"
	
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	fmt.Println("4")
	fmt.Println("Generated HOST:",os.Getenv("MYSQL_HOST"))
	fmt.Printf("ok")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
	)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
