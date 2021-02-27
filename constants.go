package main

import (
    "fmt"
);

func main() {

    /*
	-> Constant values must be available at compile time:
	    --> Eg: 1, 'x', "abc" are valid.
	    --> Eg: math.Sin(12) is not valid.
    */

    const myConst int = 12;
    const myConst2 int = 123;
    fmt.Printf("%v", myConst + myConst2);
}
