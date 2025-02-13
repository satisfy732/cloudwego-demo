package mysql

import (
	"fmt"
	"os"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	fmt.Println(os.Getenv("MYSQL_HOST"))
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
	os.Getenv("MYSQL_USER"),
	os.Getenv("MYSQL_PASSWORD"),
	os.Getenv("MYSQL_HOST"),
	os.Getenv("MYSQL_DATABASE"),
)
    fmt.Println("DSN:", dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.Student{})
	
	type version struct {
		Version string
	}
	var v version
	err = DB.Raw("SELECT VERSION() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql version:", v.Version)
}
