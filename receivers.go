package main

import "fmt"

type User struct {
    name string;
};

func (u *User) changeName() string {
    u.name = "asdasd";
    return u.name;
}

func main() {
    newUser := User {
	"mano",
    };

    fmt.Printf("%v", newUser);
    changeName(&newUser);
    fmt.Printf("%v", newUser->changeName());
}
