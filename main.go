package main

import "fmt"

func main() {
	var x int
	fmt.Scanln(&x)
	s := make([]int, x)
	cont := 0
	for i, _ := range s {
		fmt.Scanln(&s[i])
		cont += s[i]
	}
	fmt.Println(cont)
}
