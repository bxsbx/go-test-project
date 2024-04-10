package main

import (
	"StandardProject/migration/database/dm"
	"github.com/golang-migrate/migrate/v4"
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
	filePath := "file://" + wd + "/migration/dm/file"

	return filePath, nil
}

// dirty = 1 跳过脏文件（无法执行的文件），继续执行下面版本的文件，慎用
func main() {
	//schema=%s&compatibleMode=Mysql
	dmUrl := "dm://SYSDBA:SYSDBA@127.0.0.1:5236?schema=teachers_awards123&compatibleMode=Mysql"

	//获取文件路径
	filePath, err := migrateFilePath()
	if err != nil {
		log.Fatal(err)
	}

	p := &dm.DM{}
	dbDriver, err := p.Open(dmUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer dbDriver.Close()

	m, err := migrate.NewWithDatabaseInstance(filePath, "dm", dbDriver)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	// 更新到最新
	//if err := m.Up(); err != nil {
	//	log.Fatal(err)
	//}

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
	//err = m.Force(20230623143423)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// n > 0 ：更新 ，n < 0 ：回退 n表示步数
	//if err := m.Steps(1); err != nil {
	//	log.Fatal(err)
	//}

}
