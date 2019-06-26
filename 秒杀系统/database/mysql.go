package database

import (
	"Golang_Study/秒杀系统/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	_ "mysql"
	"os"
)

var DBsql *gorm.DB

func init() {
	var err error
	////open创建一个连接池
	                     //                  用户名:密码@tcp(主机地址:端口)/数据库名?charset=utf8&parseTime=True&loc=Local
	DBsql, err = gorm.Open(config.ChDataBase.Types, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ChDataBase.User,
		config.ChDataBase.PassWord,
		config.ChDataBase.Host,
		config.ChDataBase.DataBase))
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	DBsql.DB().SetMaxIdleConns(20)//用于设置闲置的连接数
	DBsql.DB().SetMaxOpenConns(100)//用于设置最大打开的连接数，默认值为0表示不限制。
	//设置默认表前缀
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName  string) string  {
		return config.ChDataBase.Prefix + defaultTableName
	}
	// 全局禁用表名复数
	DBsql.SingularTable(true)
}