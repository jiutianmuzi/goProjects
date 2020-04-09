package main
import "fmt"

func main(){
	fmt.Println("lixusheng")
	var a int = 2
	var b int = 3
	fmt.Println(max(a,b))
	
}

func max(num1,num2 int) int{
	if(num1 > num2){
		return num1
	}else{
		return num2
	}
}