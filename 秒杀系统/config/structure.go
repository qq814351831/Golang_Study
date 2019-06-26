package config

import "time"

type App struct {
	//分页大小
	PageSize int
}
type Database struct {
	Types string
	User string
	PassWord string
	Host string
	DataBase string
	Prefix string
}
type Server struct {
	HttpIp string
	HttpPort string
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
}

