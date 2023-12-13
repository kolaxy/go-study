package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   []string{"Moscow", "Klinton", "kirovsk"}, // города с населением 10-99 тыс. человек
		100:  []string{"CPB", "NVZ", "Briansk"},         // города с населением 100-999 тыс. человек
		1000: []string{"KIpr", "Harkov", "Novosibirsk"}, // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Moscow":     54,
		"Klinton":     65,
		"kirovsk":     789,
		"CPB":         345,
		"NVZ":         548,
		"Briansk":     579,
		"KIpr":        65468,
		"Harkov":      78645,
		"Novosibirsk": 57498,
	}

	for k := range cityPopulation {
		array := groupCity[100]
		flag := true
		for _, string := range array {
			if k == string {
				flag = false
			}
		}
		if flag {
			delete(cityPopulation, k)
		}
	}
	fmt.Println(cityPopulation)
}
