package sysinit

import (
	"IrisAdminApi/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	var (
		conn string
		err  error
	)

	if config.Config.DB.Adapter == "mysql" {
		// 执行mysql 链接
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.Name)
	} else {
		panic("不支持非mysql的链接方式")
	}

	fmt.Println(conn)
	if DB, err = gorm.Open(config.Config.DB.Adapter, conn); err != nil {
		panic(fmt.Sprintf("链接Mysql失败:%s", err.Error()))
	}

	gorm.DefaultTableNameHandler = func(Db *gorm.DB, defaultTableName string) string {
		return config.Config.DB.Prefix + defaultTableName
	}

	if config.Config.Env == "dev" {
		// 开发环境开启sql 日志
		DB.LogMode(true)
		DB.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	DB.DB().SetMaxOpenConns(50) // mysql 最大链接数
	DB.DB().SetMaxIdleConns(10) // mysql 最大限制数
}
