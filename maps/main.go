package main

import "fmt"

func main() {

	m := map[string]int{
		"jeremy": 24,
		"john":   22,
		"josh":   27,
	}

	k := keys(m)

	fmt.Printf("Map is: %v\n", m)
	fmt.Printf("Keys are: %v\n", k)
}

func keys(m map[string]int) (s []string) {

	s = make([]string, len(m))

	for name, _ := range m {
		s = append(s, name)
	}

	return s
}
