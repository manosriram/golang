package main

import "fmt"

/*
    1. Boolean (bool)
    2. Integer (int)
    3. Unsigned Integer (uint)
    4. Float (float64)
    5. String (string)
*/

func main() {
    var v1 bool = false;
    var v2 int = 128;
    var v3 uint16 = 123;
    var v4 float64 = 0.5;
    var s string = "asdlaskjda";
    byteArray := []byte(s);
    var add = v2 + int(v3);



    fmt.Printf("%v\n", v1);
    fmt.Printf("%d\n", v2);
    fmt.Printf("%d\n", v3);
    fmt.Printf("%d\n", add);
    fmt.Printf("%v\n", s);
    fmt.Printf("%v\n", byteArray);
    fmt.Printf("%v\n", 2 * v4);
}
