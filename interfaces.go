package main

import "fmt"

type Animal interface {
    Says() string;
    NumberOfLegs() int;
};

type Dog struct {
    Name string;
    breed string;
};

type Gorilla struct {
    Name string;
    Color string;
    NoOfTeeth int;
};

func main() {
    dog:= Dog {
	Name: "Baira",
	breed: "Country",
    };

    printInfo(dog);
}

func printInfo(a Animal) {
    fmt.Printf("This animal says %s and has %d no of legs.", a.Says(), a.NumberOfLegs());
};

func (d Dog) Says() string {
    return "woof";
};

func (d Dog) NumberOfLegs() int {
    return 4;
};
