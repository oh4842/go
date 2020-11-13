package main

import "fmt"

func main(){
	array1 := [5]string{"i", "n", "h", "a", "!"}
	array3 := [5]string{"a","b","c","d","e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	fmt.Println(slice3, slice4)
	//array3[2] = "z"
	slice1 := array1[0:3]
	fmt.Println(slice3, slice4)
	fmt.Println(slice1)

	slice3[2] := "Queen"
	fmt.Println(slice3)
}
