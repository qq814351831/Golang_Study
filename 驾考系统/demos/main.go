package main

import (
	"fmt"
	"sync"
	"time"
	"utils"
)
/*
考场签到,名字丢入管道
只有五个车道,最多供5个人同时考试
考生按签到顺序依次考试,给予考生10%的违规几率
每3秒钟巡视依次,发现违章的清出考场,否则输出考场时序良好
所有考试者考完后,生成考试记录
当前目录下的成绩录入mysql数据库,数据库允许一写多读
成绩录入完毕通知考生,考生查阅自己的成绩
*/
//定义一个全局的大管道放考生名字
var (
	//计数器
	wg sync.WaitGroup
)

func main() {
	for i := 0;i<20;i++  {
		chNames <- utils.GetRandomName()//考生名字丢入管道
	}
	//关闭管道
	close(chNames)
	/*巡考*/
	go Patrol()
	/*考生并发考试*/
	for name := range chNames{
		//fmt.Println(name)
		wg.Add(1)//计数器设置1
		go func(name string) {
			TakeExam(name)//每个人起一个协程
			wg.Done()
		}(name)
	}
	wg.Wait()//阻塞
	fmt.Println("考试完毕!")

	/*录入成绩到mysql*/
	wg.Add(1)
	go func() {//匿名函数
		utils.WriteScore2Mysql(scoreMap)//入库
		wg.Done()
	}()
	 //故意给一秒间隔确保WriteScore2DB先抢到数据库的读写锁
	<-time.After(1*time.Second)
	/*查询自己的成绩*/
	for _,name := range examers{
		wg.Add(1)
		go func(name string) {
			QueryScore(name)
			wg.Done()
		}(name)
	}
	<-time.After(1*time.Second)
	for _,name := range examers{
		wg.Add(1)
		go func(name string) {
			QueryScore(name)
			wg.Done()
		}(name)

	}
	wg.Wait()//阻塞
	fmt.Println("END")
}




