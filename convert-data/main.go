package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	fields := strings.Split(data, ",")
	age, _ := strconv.Atoi(fields[1])
	return User{Name: fields[0], Age: age, Address: fields[2]}
}

func main() {
	data := "Edo,20,Jakarta"
	user := ConvertData(data)
	fmt.Printf("name : %s, age: %d address: %s", user.Name, user.Age, user.Address)
}
