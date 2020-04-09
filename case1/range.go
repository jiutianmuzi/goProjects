package main

import "fmt"

/*
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

*/

func main(){
	
	nums := []int{2,3,4}
	sum := 0
	for _,num:= range nums{
		sum += num
	}
	
	
	fmt.Println("sum:",sum)
	
	for i,num := range nums {
		if num == 3 {
			fmt.Println("index:",i)
		}
	}
	/*
	定义map
		var map_variable map[key_data_type]value_data_type
	*/
	
	kvs := map[string]string{"a":"apple","b":"banana"}
	
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n",k,v)
	}
	
	for i,c := range "ggoo"{
		fmt.Println(i,c)
	}
	
}