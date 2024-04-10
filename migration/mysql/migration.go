package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"log"
	"os"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migration 文件路径
func migrateFilePath() (string, error) {
	wd, err := os.Getwd()
	if nil != err {
		return "", err
	}
	wd = strings.ReplaceAll(wd, "\\", "/")
	filePath := "file://" + wd + "/migration/mysql/file"

	return filePath, nil
}

// dirty = 1 跳过脏文件（无法执行的文件），继续执行下面版本的文件，慎用
func main() {

	// 读取配置文件
	cfg, err := config.NewConfig("ini", "beego/conf/app.conf")
	if err != nil {
		log.Fatal(err)
	}

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		cfg.String("StandardProject::dbUserName"),
		cfg.String("StandardProject::dbPassword"),
		cfg.String("StandardProject::dbHost"),
		cfg.String("StandardProject::dbPort"),
		cfg.String("StandardProject::dbName"),
	)

	//获取文件路径
	filePath, err := migrateFilePath()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(filePath, "mysql", dbDriver)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	// 更新到最新
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	// 当前版本 ? 指定版本  > ：回退 ，< ： 更新，(指定的版本号需要在文件中存在)
	//if err := m.Migrate(202306231434); err != nil {
	//	log.Fatal(err)
	//}

	//version, dirty, err := m.Version()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(version, dirty)
	//强制回到某个版本号
	//err = m.Force(202306231434)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// n > 0 ：更新 ，n < 0 ：回退 n表示步数
	//if err := m.Steps(1); err != nil {
	//	log.Fatal(err)
	//}

}
