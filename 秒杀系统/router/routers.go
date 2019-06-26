package router

import (
	"Golang_Study/秒杀系统/controller"
	"github.com/gin-gonic/gin"
)

func HttpServer()*gin.Engine {
	//这是默认的服务器。使用gin的Default方法创建一个路由Handler；
	r := gin.Default()

	/*然后通过Http方法绑定路由规则和路由函数。不同于net/http库的路由函数，gin进行了封装，把request和response都封装到了gin.Context的上下文环境中*/
	r.GET("/login",controller.Login)
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "ling",
	//	})
	//})
	return r
}
