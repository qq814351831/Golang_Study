/*
1 接口?
不是传统的面向对象的概念 没有类和继承的概念
接口: 一组方法的集合
2 接口的定义?
type st_name struct{

}
type i_name interface{
	method1(param_list)return_ist
	method1(param_list)return_ist
	......
}
3 调用接口?
var I_name I_ST
var ST_name ST_
I_name = new(ST_name)
ST_name要求必须实现I_ST的方法
4 接口的拓展
func func_name(i interface{}){

}
*/

package main

import "fmt"

//定义接口
type bb interface {
	GET()
	POST()
}
//定义结构体
type aa struct{

}
//实现方法
func (this *aa)GET(){
}
func (this *aa)POST(){
}

func dd(i interface{}){
	_=i
	fmt.Println(i)
}

func main(){
	//调用
	var cc bb// 定义一个接口类型的变量
	cc = new(aa)
	_=cc
	cc.GET()
	cc.POST()

	dd(1020)
	dd("www.cwy.com")
}
