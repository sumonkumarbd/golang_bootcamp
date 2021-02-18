package main

import "fmt"

// write programe codes here...

// Day=01 --lesson of run first golang code..
func first_go() {

	fmt.Println("hello world")
}
//end of first_go ()

//start a function to test arguments
func add_num(a , b int) int{
   // return a + b
   // return a - b
   // return a * b
    return a / b
}

// main function
func main() {
	// This is main function note::'all the programe will be able to exiqute in this main function.'

	// to run day=01 programe
	//first_go()
	
	//run function add_num
	fmt.Println(add_num(10,5))
}
