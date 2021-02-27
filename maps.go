package main

import ("fmt"
    );

func main() {
    emails := make(map[string]string);

    emails["mano"] = "mano.sriram0@gmail.com";
    fmt.Println(emails["mano"]);
    delete(emails, "mano");
    is := emails["mano"];

    if (len(is) != 0) {
	fmt.Println("mano value is present");
    }
}
