package main

import "fmt"

/*
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，
功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)

*/

func main(){
	
	slice1 := make([]int ,10)
	i := 0
	
	for ; i< 10;i++ {
		slice1[i] = i+100
	}	
	
	fmt.Println(slice1)
	
	var numbers = make([]int,3,5)
	printSlice(numbers)
	
	
	
}

func printSlice(x []int){
	fmt.Printf("len=%d cap%d slice=%v\n",len(x),cap(x),x)
}


