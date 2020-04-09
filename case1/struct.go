package main

import "fmt"

/*

type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

*/

type Books struct {
	
	title string
	author string
	subject string
	book_id int
	
}

func main(){
	fmt.Println(Books{title :"Go 语言",author: "Thompson"})
	
	var Book1 Books
	Book1.author = "bule tiger"
	
	fmt.Printf(Book1.author)

}