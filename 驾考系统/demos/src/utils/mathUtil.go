package utils

import (
	"math/rand"
	"sync"
	"time"
)
var(
	//随机数互斥锁(确保随机数GetRandomInt不能被并发访问)
	ramdomMUtex sync.Mutex
)
//定义一个函数,生成随机数
func GetRandomInt(start, end int)int  {
	// time.Millisecond    表示1毫秒
	// time.Microsecond    表示1微妙
	// time.Nanosecond    表示1纳秒
	//time.Sleep()          休眠多少时间,处于阻塞状态,后续程序无法执行
	//time.After()          非阻塞,可用于延迟
	//time.Since()          两个时间点的间隔
	ramdomMUtex.Lock()//互斥锁加锁
	<-time.After(1*time.Nanosecond)//延迟一秒
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start+1)
	ramdomMUtex.Unlock()//互斥锁解锁
	return n
}
