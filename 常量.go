package main //命名空间

import "fmt" //导入包

func main() {//main方法相当于入口   项目必须有
	fmt.Println("hello")
	//iota 是特殊的常量  可以是一个被编译器修改的常量
	//特点: 每一个const关键字出现时,被重置为0(const 内部的第一行之前),const 中每新增一行常量声明将使 iota 计数一次

	//定义常量
	const LENGTH int = 10
	const WIDE int = 5
	var aaa int

	aaa = LENGTH * WIDE
	fmt.Println("面积为:", aaa)

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)


}