/*
条件语句
if else
特殊写法
if _,error:=func();error !=nil{

}else{

}

*/

package main

import "fmt"

var 中文 int
//init函数用来初始化变量会执行在main函数之前
func init(){
	中文 = 888
}

func main(){
	if 中文 == 999{

	}else{
		fmt.Println("这是中文",中文)
	}
}
