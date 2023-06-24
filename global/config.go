package global

import (
	"StandardProject/common/gormdb"
	"github.com/astaxie/beego/config"
	"log"
)

// 全局配置初始化
func init() {

	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatal(err)
	}
	//logz.LogConfig(cfg)
	//tracer.Config(cfg)
	gormdb.InitDB(cfg)
	//redis.InitRedis(cfg)
	//mongodb.InitMongoDB(cfg)
	//mq.InitMq(cfg)

}
