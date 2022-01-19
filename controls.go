package main

import "fmt"

/*
1. if-else statement
    if(condition){
		//statements
	}

	if(condition){
		//statement
	}
	else{
		//statement
	}

	if(condition){
		//statement
	}
	else if{
		// statement
	}
	else{
        //statement
	}


2. switch(data){
case var1:
	  statment
case var2:
	 statement
case var3:
	statement
default:
	statement

}
*/
func main() {
	// fmt.Print("Enter a number:")
	// var input int
	// fmt.Scanln(&input)

	// if input%2 == 0 {
	// 	fmt.Println("Hey you are even")
	// } else {
	// 	fmt.Println("Hey you are odd")
	// }
	// x := 1000
	// if x := 10; x%2 == 0 {
	// 	fmt.Println("Hey you are even")
	// } else {
	// 	fmt.Println("Hey you are odd")
	// }
	// fmt.Println(x)

	data := 10
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
		fallthrough //exrcute the next case also
	case 100:
		fmt.Println(data)
	case 103:
		fmt.Println(data)
	case 109:
		fmt.Println(data)
	case 102:
		fmt.Println(data)
	case 15:
		fmt.Println(data)
	default:
		data = 10909
		fmt.Println(data)
	}
}
