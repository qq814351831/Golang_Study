//数组定义
//var 变量名 = [数组长度]变量类型{数组值}

package main

import "fmt"

var aa = [20]string{"1","2","3"}

func main(){
	for i:=0;i<len(aa);i++{
		fmt.Println(aa[i])
	}

}
