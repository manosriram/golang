package main

import "fmt"

func main() {
    /*
	Array Format:
	
	var arrayName [size]dataType;
    */

    var a [2]int;  // Defaults to 0 for each array element
    var b [2]string;

    anotherArray := [2]int{0, 23}; // Declare and assign.

    fmt.Println(a, b);
    fmt.Println(anotherArray);
}
