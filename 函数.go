package main

import "fmt"

func main(){
	res :=test(10)
	fmt.Println(res)
}
//func 函数名(参数变量 参数类型)返回类型{
// 函数体
// return 返回列表
// }
func test(a int)int{
	return a
}
