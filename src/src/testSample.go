package main

import (
 "fmt"
)

const (
	a=iota
	B
	C
)

type StringChild string

//First letter in Caps means public and small means private
type Salutation struct {
	name string
	message string
}

func main() {
	
	message:= "Hello Go World!!"
	
	var anotherMessage StringChild = "Hello Another Go World !"
	
	pointerMsg := &message;//Assign memory address to the pointer
	//Other ways to achieve the same
	//var pointerMsg = &message
	//var pointerMsg *string = &message 
	
	fmt.Println(message, anotherMessage, pointerMsg, *pointerMsg, a,B,C)
	
	var s = Salutation{"Adithya", "Hola!"}
//	var s1 = Salutation{message: "Hola!", name: "Adithya"}
//	var s2 = Salutation{} s.name="kuch bhi"

	fmt.Println(s.name, s.message, s1)		
		
}

