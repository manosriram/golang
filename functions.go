package main

import "fmt"

/*
    function format:
	
    func nameOfTheFunction(arg1 arg1Datatype, arg2 arg2Datatype) returnType {
	...
    };
*/

func helloWorld(name string) string {
    return "Hello, " + name;
};

func main() {
    fmt.Println(helloWorld("Mano"));
}
