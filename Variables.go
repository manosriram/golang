package main;

import "fmt";

func main() {

    // := is used when there is no information about the variable.

    /*
	1. var foo int;
	2. var foo int = 42;
	3. foo := 42;
    */

    var a int = 12;
    fmt.Printf("%d", a);
    fmt.Printf("%T", a); // Type of 'a'
}
