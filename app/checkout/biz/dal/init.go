package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/checkout/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
