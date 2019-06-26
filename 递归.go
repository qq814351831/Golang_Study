/*
递归函数
1 什么是递归函数?
  程序在运行的时候自己调用自己\
2 规则
  func data(){
	data()
  }
3 阶乘的规则
  n! = n*(n-1)
*/

package main

import "fmt"

func main(){
	var i uint64 = 15
	ret := data(i)
	fmt.Println("%d 的阶乘是 %d",i,ret)
}
func data(n uint64)(ret uint64){
	if n>0{
		ret = n*data(n-1)
		return ret
	}
	return 1
}