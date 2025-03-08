package main

import (
	"fmt"
	"syntax-cafe-go/class_1/homework"
)

func main() {
	fmt.Printf("Enter a number: ")
	var i int 
	fmt.Scanf("%v",&i)

	if(i%2 == 0){
		homework.MainFunc()
	}else{
		fmt.Println("Odd number ")
	}

}
