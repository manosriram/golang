package main

import (
    "log"
);

func calculateValue(intChan chan int) {
    intChan <- 34343;
};

func main() {
    intChan := make(chan int);
    defer close(intChan);

    go calculateValue(intChan);

    num := <-intChan;
    log.Println(num);
}
