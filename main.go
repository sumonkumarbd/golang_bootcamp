package main

import "fmt"

// write programe codes here...

// Day=01 --lesson of run first golang code..
func first_go() {

  fmt.Println("hello world with golang")
}


//pointer examples
func pointer(){
var num *int
var num1 int = 5
    num = &num1

fmt.Println(" first declear pointer variable is :",num ,"\n","num1 variable :", num1 ,"\n","num1 variable memory number is : " , num,"\n","After de-raffarance num value is : ",*num)


}
//end of pointer example



// main function
func main() {
  // This is main function note::'all the programe will be able to exiqute in this main function.'

  // to run day=01 programe
  // first_go()
  
  //run pointer function
   pointer()
}
