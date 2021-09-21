package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type userconfig struct {
	IPaddr	string
	User	string
	Pword	string
}

func main() {
	file, _ := os.Open("test_conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := userconfig{}
	fmt.Println(conf)

	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf)
}
