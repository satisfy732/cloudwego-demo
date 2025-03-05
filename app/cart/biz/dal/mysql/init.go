package mysql

import (
	"os"
	"fmt"
	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, 
	os.Getenv("MYSQL_USER"), 
	os.Getenv("MYSQL_PASSWORD"), 
	os.Getenv("MYSQL_HOST"))
	

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		//nolint:errcheck
		DB.AutoMigrate(
			&model.Cart{},
		)
	}
}
