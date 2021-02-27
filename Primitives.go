package main

import "fmt"

/*
    1. Boolean (bool)
    2. Integer
	- int
	- int8
	- int16
	- int32
	- int64
    3. Unsigned Integer
	- uint
	- uint8
	- uint16
	- uint32
	- uint64
	- uintptr
    4. Float
	- float32
	- float64
    5. String (string)
    6. Complex
	- complex64
	- complex128
*/

func main() {
    var v1 bool = false;
    var v2 int = 128;
    var v3 uint16 = 123;
    var add = v2 + int(v3);
    var v4 float64 = 0.5;
    var s string = "asdlaskjda";
    byteArray := []byte(s);

    fmt.Printf("%v\n", v1);
    fmt.Printf("%d\n", v2);
    fmt.Printf("%d\n", v3);
    fmt.Printf("%d\n", add);
    fmt.Printf("%v\n", s);
    fmt.Printf("%v\n", byteArray);
    fmt.Printf("%v\n", 2 * v4);
}
