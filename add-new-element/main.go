package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		n := make([]int, len(data)+1)
		n[0] = newData
		copy(n[1:], data)
		return n
	}
	if position == "down" {
		return append(data, newData)
	} else {
		return nil
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	new := 6
	post := "up"

	newArray := AddElement(data, new, post)
	fmt.Println(newArray)

	post = "down"
	newArray = AddElement(data, new, post)
	fmt.Println(newArray)
}
