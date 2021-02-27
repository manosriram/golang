package main

import "fmt"

func main() {
    array := [5]int{100, 200, 300, 400, 500};
    fmt.Println(array);

    i := 0;
    for (i < len(array)) {
	fmt.Println(array[i]);
	i++;
    }
    fmt.Println();

    for t:=1;t<len(array);t++ {
	fmt.Println(array[t]);
    }

}
