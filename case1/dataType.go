package main
import "fmt"

func main(){
	/*
	变量声明
	第一种
	 var identifier type
	 var identifier1, identifier2 type
	
	第二种
	根据值自行判定变量类型。
	var v_name = value
	
	第三种
	:= 左侧如果没有声明新的变量，就产生编译错误
	v_name := value
	
	*/
	
	var a bool = true
	
	var b int = 100
	
	var c float32 = 100
	
	var d string = "Runoob"
	
	fmt.Println(a);
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	
	intVal := -1;
	fmt.Println(intVal)
	
	/*
	常量
	const identifier [type] = value
	*/
	const dd = "abc"
	fmt.Println(b)
	
	
}