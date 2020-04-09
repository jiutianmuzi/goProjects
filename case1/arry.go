package main

import "fmt"

func main(){
	
	/*
	var variable_name [SIZE] variable_type
	
	*/
	
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0} 
	
	var i int = 0
	for ;i<5;i++ {
		fmt.Printf("Element[%d] = %f\n",i,balance[i])
	}
	
}