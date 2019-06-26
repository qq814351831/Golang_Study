package main

import (
	"./config"
	"./router"
)

func main() {
	Server := config.ChServer.HttpIp + ":" + config.ChServer.HttpPort
	//路由
	r := router.HttpServer()
	//最后启动路由的Run方法监听端口。还可以用http.ListenAndServe(":8080", router)，或者自定义Http服务器配置。
	r.Run(Server) // listen and serve on 0.0.0.0:8080

}
