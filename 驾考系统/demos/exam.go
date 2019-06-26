package main

import (
	"fmt"
	"time"
	"utils"
)
var(
	//所有考生
	//        make(参数类型,分配空间,预留空间)
	chNames = make(chan string,100)
	//考试者的名字
	examers = make([]string,0)
	//5条车道管道
	chLanes = make(chan int,5)
	//违纪管道
	chFouls = make(chan string,100)
	//考试成绩
	scoreMap = make(map[string]int)
)



//巡考
func Patrol(){
	ticker := time.NewTicker(1 * time.Second)
	for {
		//fmt.Println("老王正在巡考...")
		//没有人违纪就读不出来
		select {
		case name := <-chFouls:
			fmt.Println(name,"考试违纪!!!")
		default:
			fmt.Println("考场秩序良好")
		}
		<-ticker.C
	}
}

//考试
func TakeExam(name string){
	chLanes <- 123//进入管道
	fmt.Println(name,"正在考试...")//考生依次考试,每次五个
	//将参与考试的考生姓名记录
	examers = append(examers, name)
	//生成考试成绩  这里因为是每次五个一起访问 会产生并发  得到的值是一样的
	score := utils.GetRandomInt(0,100)
	scoreMap[name] = score
	if score < 10 {
		chFouls <- name//违纪人丢到违纪管道中,考试过程中这个管道不能关闭它
		//fmt.Println(name,"考试违纪!!!",score)
	}
	//考试持续5秒钟
	//<-time.After(2*time.Second)
	<-time.After(400*time.Millisecond)//巡视五次
	<-chLanes//踢出管道
	//wg.Done()//计数器-1


}

//二级缓存查询成绩
func QueryScore(name string){
	//首先查询redis
	score,err := utils.QueryScoreFromRedis(name)
	if err != nil {
		//score,_ = utils.QueryScoreFromMysql(name)
		argsMap := make(map[string]interface{})
		argsMap["name"] = name
		scores := make([]utils.ExamScore,0)
		err = utils.QueryFromMysql("score",argsMap,&scores)//查询mysql
		utils.HandlerError(err,`utils.QueryFromMysql("score",argsMap,&scores)`)

		fmt.Println("mysql成绩",name,":",scores[0].Score)
		//将数据写入redis
		utils.WriteScore2Redis(name,scores[0].Score)
	}else{
		fmt.Println("redis成绩",name,":",score)
	}

}
