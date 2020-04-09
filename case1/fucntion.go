package main

import "fmt"

func main(){
	
	var str string = "lixusheng"
	var num1 int = 100
	var num2 int = 1
	
	funcOne(str,num1)
	
	fmt.Println(funcOne(str,num2))

}

/*
	func function_name( [parameter list] ) [return_types] {
		函数体
	}
*/


func funcOne( str string,num int ) string {

	if num >2 {
		fmt.Println("funcOne executed")
	}
	
	return str
}