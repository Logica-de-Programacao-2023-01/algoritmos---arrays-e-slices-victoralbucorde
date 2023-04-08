package main

import "fmt"

func main() {
	var i int
	conj := [5]int{0, 10, 2, 3, 4}
	var crescente bool

	fmt.Println("Seu conjunto é:", conj)
	for len(conj) > i+1 {
		if conj[i+1] > conj[i] {
			crescente = true
		} else {
			crescente = false
			break
		}
		i++
	}
	if crescente == true {
		fmt.Println("Seu conjento é crescente.")
	} else {
		fmt.Println("Seu conjunto não é crescente.")
	}
}
