package main

import "fmt"

func main() {
	var data int //declaring variable
	data = 10    //assigning variable
	fmt.Println(data)

	data1 := 1000 //it will declare the variable and assign value to it
	fmt.Println(data1)

	//Node->If we declare a variable to int type (data1 :=10) we cannot assign another type data later

	// data := 100
	// fmt.Println(data) // it will gives us error redclaring

	//data type
	//1.Primitive
	//int,int16,int32,int32,uint,float,float64,string,bool,complex,byte

	//2.Non-primitive
	//strcuct,map,chan,interface,array,slice,rune,reflect,pointer

	//functional programming
	//recent feature update of java->from java 11

	name := "rahul"
	fmt.Println(name)
}
