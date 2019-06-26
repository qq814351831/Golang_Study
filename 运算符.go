/*
Go运算符
1.算数运算符
+ - * /
2.关系运算符
== != > < >= <=
3.逻辑运算符
&& || !
4.位运算符
& | ^
5.赋值运算符
= += -= /= %=
6.其他运算符
& *
*/
package main

import "fmt"

func main(){
	var a = 21
	var b = 10
	var c int
	c = a+b
	fmt.Println("a+b=",c)
	c = a-b
	fmt.Println("a-b=",c)
	if(a > b){
		fmt.Println("a>b")
	}else{
		fmt.Println("a<=b")
	}

	var d = true
	var f = false
	if(d&&f){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}
