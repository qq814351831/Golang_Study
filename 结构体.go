//语言结构体
//结构体将另个或者多个任意类型的命名变量组合在一起的聚合数据类型。
package main

import "fmt"

//定义结构体
//type 名字 struct{
//  name type
//  .....
// }
/*
结构体指针
type name *struct_name
*/

type student struct {
	ID int
	Url string
}
func main(){
   //结构体赋值
   var stu student


   stu.ID = 101
   stu.Url = "www.CWY.com"

   fmt.Println(stu)
   //当传值的时候,相当于copy一份副本传到Run里面,原数据不会做任何改变
   //Run(stu)
   //fmt.Println(stu)
   /*
   	{101 www.CWY.com}
	{101 www.CWY.com}
   */
	//当传地址的时候,改变的是原数据地址对应的值
	Run(&stu)
	fmt.Println(stu)
	/*
	{101 www.CWY.com}
	{111 www.CWY.com}
	*/
}
//函数调用
func Run(Stu *student)(data interface{}){
	//fmt.Println(Stu.ID)
	Stu.ID = 111
	return Stu
}
