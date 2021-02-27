package main

import "fmt"

func main() {
    ids := []int { 123, 1234, 12345 };
    for _, id := range ids {
	fmt.Println(id);
    }

    var emails = make(map[string]string);
    emails["mano"] = "mano.sriram0@gmail.com";
    for key, value := range emails {
	fmt.Printf("%s - %s", key, value);
    }
}
