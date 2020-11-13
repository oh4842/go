package main

import "fmt"

func main(){
	array3 := [5]string{"a","b","c","d","e"}
	var slice3, slice4 = []string
	//slice3 := array3[0:3]
	//slice4 := array3[2:5]
	fmt.Println(slice3, slice4)
	//slice3[2] := "Queen"
	//fmt.Println(slice3, slice4)
	fmt.Println(array3)
}
