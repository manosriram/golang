package main

import (
    "fmt"
    "log"
    "encoding/json"
)

type User struct {
    FirstName string `json:"first_name"`
    age int `json:"age"`
};

func main() {
    myJson := `[
    {
	"first_name": "Mano",
	"age": 21
    },
    {
	"first_name": "Sriram",
	"age": 52
    }]`;

    var unmarshalled []User;

    err:= json.Unmarshal([]byte(myJson), &unmarshalled);

    if (err != nil) {
	log.Printf("%v", err);
    } else {
	var slice []User;
	m1:= User {
	    FirstName: "mano",
	    age: 43,
	};

	slice = append(slice, m1);
	m2:= User {
	    FirstName : "sriram",
	    age: 22,
	};
	slice = append(slice, m2);

	newJson, err:= json.MarshalIndent(slice, "", "  ");
	if (err != nil) {
	    log.Println(err);
	}

	fmt.Println(string(newJson));
    }

}
