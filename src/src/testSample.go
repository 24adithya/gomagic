package main

import (
 "fmt"
)

func main() {
	
	message:= "Hello Go World!!"
	
	pointerMsg := &message;//Assign memory address to the pointer
	//Other ways to achieve the same
	//var pointerMsg = &message
	//var pointerMsg *string = &message 
	
	fmt.Println(message, pointerMsg, *pointerMsg)
}

