package utils

import (
	"errors"
	"fmt"
	"sync"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx" //sqlx包 database/sql的扩展
	/*MySQL*/
	_ "mysql" //mysql包

)
var(
	//数据库读写锁
	dbMutext sync.RWMutex
)

//通用查询mysql工具
                   //表名             查询条件                       接收地址
func QueryFromMysql(tableName string,argsMap map[string]interface{},dest interface{})(err error){
	//fmt.Println("Mysql")
	dbMutext.RLock()//读锁 当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景
	db, err :=sqlx.Connect("mysql","root:root@tcp(localhost:3306)/driving_exam")//连接数据库
	//数据库类型,数据库名:密码@tcp(ip:端口)/表名
	HandlerError(err,`sqlx.Connect("mysql","root:root@tcp(localhost:3306)/driving_exam")`)//记录错误
	defer db.Close()//关闭db

	selection := ""//键
	values := make([]interface{},0)
	for key,val := range argsMap{
		selection += (" and "+key+"=?")
		values = append(values,val)
	}
	selection = selection[4:]
	sql := "select * from "+tableName+" where"+selection;

	//                              加三个点将切片展开为不定长参数
	err = db.Select(dest,sql,values...)//查询成绩
	if err != nil{
		fmt.Println(err,`db.Select(&ExamScores,"select * from score where name=?;", name)`)
		return
	}
	dbMutext.RUnlock()//读锁解锁

	return nil

}

//mysql查询
//func QueryScoreFromMysql(name string)(score int,err error){
//	//fmt.Println("Mysql")
//	dbMutext.RLock()//读锁 当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景
//	db, err :=sqlx.Connect("mysql","root:root@tcp(localhost:3306)/driving_exam")//连接数据库
//	//数据库类型,数据库名:密码@tcp(ip:端口)/表名
//	HandlerError(err,`sqlx.Connect("mysql","root:root@tcp(localhost:3306)/driving_exam")`)//记录错误
//	defer db.Close()//关闭db
//
//	ExamScores := make([]ExamScore,0)
//	err = db.Select(&ExamScores,"select * from score where name = ?;",name)//查询成绩
//	if err != nil{
//		fmt.Println(err,`db.Select(&ExamScores,"select * from score where name=?;", name)`)
//		return
//	}
//	fmt.Println(ExamScores)
//
//	dbMutext.RUnlock()//读锁解锁
//
//	return ExamScores[0].Score,nil
//
//}
//redis查询
func QueryScoreFromRedis(name string)(score int,err error){
	//fmt.Println("Redis")
	conn, err  := redis.Dial("tcp","localhost:6379")//连接redis
	HandlerError(err,`redis.Dial("tcp","localhost:6379")`)
	defer conn.Close()

	reply,e := conn.Do("get",name)//请求
	if reply != nil{
		score,e = redis.Int(reply,e)//转换成int型
		//fmt.Println("!!!!!!!!!!!!",score,e)
	}else{
		return 0,errors.New("未能从redis中查到数据")
	}

	if err != nil{
		fmt.Println(err,`conn.Do("get",name)或者redis.Int(reply,err)`)
		return 0,e
	}

	return score,nil
}
//mysql录入成绩
func WriteScore2Mysql(scoreMap map[string]int){
	//锁定为写模式,期间不允许读
	dbMutext.Lock()//加锁
	//                  数据库类型
	db, error :=sqlx.Connect("mysql","root:root@tcp(localhost:3306)/driving_exam")//连接数据库
	//数据库类型,数据库名:密码@tcp(ip:端口)/表名
	HandlerError(error,`sqlx.Connect("mysql","root:root@tcp(localhost:3306)/driving_exam")`)//记录错误
	defer db.Close()//关闭db
	for name,score :=range scoreMap{
		_, err := db.Exec("insert into score (name,score) values( ?, ?);", name, score)//新增数据
		HandlerError(err,`db.Exec("insert into score (name,score) values( ?, ?);", name, score)`)//记录错误
		fmt.Println("插入成功!")
	}
	fmt.Println("成绩录入完毕!!!")
	dbMutext.Unlock()//解锁,开放查询

}
//redis录入成绩
func WriteScore2Redis(name string,score int)(err error){
	conn, err  := redis.Dial("tcp","localhost:6379")//连接redis
	HandlerError(err,`redis.Dial("tcp","localhost:6379")`)
	defer conn.Close()

	_, err = conn.Do("set",name,score)
	HandlerError(err,`conn.Do("set",name,score)`)

	fmt.Println("写入redis成功!!!")
	return err

}
