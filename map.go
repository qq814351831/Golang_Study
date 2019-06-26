/*
1 map概念?
  本身是key=>value的集合
  特点: 无序的键值对的集合
2 定义map
  var name map[key_type]value_type
  var Golang map[int]string
3 map值的存储
一定要初始化 map
4 删除map的数据
 delete()
5 range 遍历 map数据

*/
package main

import "fmt"

var Golang = make(map[int]string)

func main(){
  Golang[1]="www.cwy.com"
  fmt.Println(Golang)
  delete(Golang,1)
  fmt.Println(Golang)
}