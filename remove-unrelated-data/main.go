package main

import (
	"fmt"
	"strings"
)

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	keys := strings.Split(key, ",")
	for _, n := range keys {
		delete(dataMap, n)
	}
	//your code here

	return dataMap
}

func main() {
	data1 := map[string]any{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	fmt.Println(removeUnrelated(data1, "address"))

	data2 := map[string]any{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	fmt.Println(removeUnrelated(data2, "jump"))

	data3 := map[string]any{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	fmt.Println(removeUnrelated(data3, "tiga"))
}
