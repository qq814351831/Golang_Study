package main

import "fmt"

//切片(slice)
/*
1 slice定义?
对于数组的抽象: 数组满足不了需求,申请空间不够大的时候,使用slice
2 slice定义
	var name []type
	[1]
	var Golang []int
	[2]
	var name []type=make([]type,len)
3 slice初始化
	Golang :=[]int{1,2,3}
4 内置函数
	len() append()
*/

func main(){
	Golang :=[]int{1,2,3}
 	fmt.Println(len(Golang))//长度
	Golang = append(Golang,4)//添加
	fmt.Println(Golang)
}
