package global

import (
	mongodb "StandardProject/sync/ebag/mongdb"
	"StandardProject/sync/ebag/mysql"
	"StandardProject/sync/ebag/redis"
	"github.com/astaxie/beego/config"
	"log"
)

type Config struct {
	ZsDomain     string
	DeviceId     string
	Appsec       string
	Paper2Domain string
	Paper2Token  string

	TeachingAdminDomain string
	TeachingAdminToken  string
}

var (
	MyConfig Config
)

func init() {
	cfg, err := config.NewConfig("ini", "sync/ebag/conf/app.conf")
	//cfg, err := config.NewConfig("ini", "conf/main.conf")
	if err != nil {
		log.Fatal(err)
	}
	MyConfig.ZsDomain = cfg.String("SyncConfig::ZsDomain")
	MyConfig.DeviceId = cfg.String("SyncConfig::DeviceId")
	MyConfig.Appsec = cfg.String("SyncConfig::Appsec")
	MyConfig.Paper2Domain = cfg.String("SyncConfig::Paper2Domain")
	MyConfig.Paper2Token = cfg.String("SyncConfig::Paper2Token")

	MyConfig.TeachingAdminDomain = cfg.String("SyncConfig::TeachingAdminDomain")
	MyConfig.TeachingAdminToken = cfg.String("SyncConfig::TeachingAdminToken")

	redis.InitRedis(cfg)
	mongodb.InitMongoDB(cfg)
	mysql.InitDB(cfg)
}
