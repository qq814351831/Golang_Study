//变量的作用域
package main

import "fmt"

var b int = 10

func main(){
//局部变量
var a,c int = 0,100
fmt.Println(a,b)
ret :=sum(a,c)
fmt.Println(ret)
}

func sum(a,b int)int{
	fmt.Println(a,b)
	return a+b
}
