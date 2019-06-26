package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

var (
	Cfg *ini.File
	ChServer *Server
	ChDataBase *Database
	ChApp *App
)

func init(){
	var err error
	ChServer = &Server{}
	ChDataBase = &Database{}
	ChApp = &App{}
	//加载ini文件
	Cfg, err = ini.Load("./config/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config.ini': %v",err)
		os.Exit(1)
	}
	LoadAppConfig()
	LoadDataBaseConfig()
	LoadServerConfig()
}
func LoadAppConfig()  {
	sec, err :=Cfg.GetSection("app")
	if err != nil{
		log.Fatalf("Fail to parse 'config/config.ini': %v",err)
		os.Exit(1)
	}
	ChApp.PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
//数据库配置
func LoadDataBaseConfig()  {
	sec, err :=Cfg.GetSection("database")
	if err != nil{
		log.Fatalf("Fail to parse 'config/config.ini': %v",err)
		os.Exit(1)
	}
	ChDataBase.Types = sec.Key("TYPES").MustString("mysql")//数据库
	ChDataBase.User = sec.Key("USER").MustString("root")//用户名
	ChDataBase.PassWord = sec.Key("PASSWORD").MustString("root")//密码
	ChDataBase.Host = sec.Key("HOST").MustString("127.0.0.1")//数据库ip
	ChDataBase.DataBase = sec.Key("DATABASE").MustString("gin")//数据库名
	ChDataBase.Prefix = sec.Key("PREFIX").MustString("gin_")//数据表前缀
}
//服务配置
func LoadServerConfig()  {
	sec, err :=Cfg.GetSection("server")
	if err != nil{
		log.Fatalf("Fail to parse 'config/config.ini': %v",err)
		os.Exit(1)
	}
	ChServer.HttpIp = sec.Key("HTTP_IP").MustString("0.0.0.0")//服务器ip
	ChServer.HttpPort = sec.Key("HTTP_PORT").MustString("8080")//端口号
	ChServer.ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60))*time.Second//请求时间
	ChServer.WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60))*time.Second//超时时间
}
