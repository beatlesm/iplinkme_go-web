package main
 
 import (
	"fmt"
)
 func f(msg string) {
 	fmt.Println(msg)
 }

 func main() {
 	f("how")
 	//go f("to")
 	//go f("goroutine")
 	/*
 	go func (msg string) {
 		fmt.Println(msg)
 	}("going")
 	*/
 }