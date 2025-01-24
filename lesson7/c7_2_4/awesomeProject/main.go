package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

func main() {
	b := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
