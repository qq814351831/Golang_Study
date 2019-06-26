/*
func error interface{
	Error()string
}
写法
errors.New("error")
*/

package main

import "fmt"

//1.定义一个结构
type aaerror struct {
	aa int
	bb int
}
//2.实现一个error接口
//这里this代表aaerror 绑定Error方法
func(this *aaerror)Error() string{
	strMsg := "www.cwy.com"//声明变量并赋值
	return fmt.Sprintf(strMsg,this.aa)
}
//3函数 int类型 除法运算
//   函数名(参数)(返回参数)
func cc(dd int,ee int)(result int,error string){
	if ee == 0{
		Data :=aaerror{
			dd,
			ee,
		}
		error = Data.Error()
		return
	}else{
		return dd/ee,""
	}
}
func main(){
	//正常情况
	//  接收返回值    传递参数            ;判断返回值的结果
	if result,error :=cc(100,10);error == ""{
		fmt.Println("100/10 = ",result)
	}
	//错误情况
	//  接收返回值    传递参数
	if _,error :=cc(100,0);error != ""{
		fmt.Println("error is = ",error)
	}
}