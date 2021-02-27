package main

import "fmt"

func main() {
    var p *int;
    x := 4;
    ptr := &x;
    *ptr += 1; // Nothing happens to ptr, Pointer arithmetic doesn't happen in Go.
    p = &x;
    fmt.Println(ptr, p, &x);
}
