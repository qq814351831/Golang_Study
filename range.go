/*
1 range 通过for 来循环遍历集合 array slice channel map

*/


package main

import "fmt"

func main(){

 	nums :=[]int{1,2,3}
	for _, num := range nums{
		fmt.Println(num)
	}

	//map
	kvs := map[string]string{"a":"www","b":"golang"}
	for _, num :=range kvs {
		fmt.Println(num)
	}
}
