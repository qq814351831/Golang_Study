/*
指针
&是取地址符号   放到变量前面
*是取内存地址的数据
指针就是  一个指针变量指向了一个值的内存地址

*/
package main

import "fmt"

//数据表字段
type ExamScore struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Score int `db:"score"`
}
func main(){
	var aa int  = 10//声明实际变量
	var bb *int//声明指针变量
	fmt.Println("变量地址为: %x",&aa)
	bb = &aa//指针变量的存储地址
	fmt.Println("变量的地址为: %x",bb)
	fmt.Println("变量的地址为: %d",*bb)//取值
	//data := ExamScore{}
	//fmt.Println(data)
	//res := &data
	//fmt.Println(res)
}
