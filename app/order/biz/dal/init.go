package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
