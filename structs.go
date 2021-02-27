package main

import (
    "fmt"
);

/*
    Format of Struct:

    type nameOfStruct struct {
	nameOfVariable typeOfVariable;
	nameOfVariable typeOfVariable;
    };
*/

type point struct {
    x int;
    y int;
};

func main() {
    pointOne := point{1, 2}; // without pointer
    var pointTwo *point = &point{3, 4};
    fmt.Printf("%d, %d", pointOne.x, pointOne.y);
    fmt.Println();
    fmt.Printf("%d, %d", pointTwo.x, pointTwo.y);
}
