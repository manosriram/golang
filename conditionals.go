package main

import "fmt"

func main() {
    x := 5;
    y := 10;
    str := "hello";

    // If else
    if (x < y) {
	fmt.Println("up");
    } else {
	fmt.Println("down");
    }

    if (str == "hello") {
	fmt.Println("equals hello");
    }

    // Switch
    switch (x) {
    case 1:
	fmt.Println("x equals 1");
	break;
    case 5:
	fmt.Println("x equals 5");
	break;
    }

}
