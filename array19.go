package main

import "fmt"

func main() {
	var i int
	conj1 := [3]int{5, 20, 42}
	conj2 := [3]int{6, 8, 11}
	conjsoma := [3]int{}

	fmt.Println("Seus conjuntos são:")
	fmt.Println(conj1)
	fmt.Println(conj2)

	for i <= 2 {
		conjsoma[i] = conj1[i] + conj2[i]
		i++
	}
	fmt.Println("Sua soma é: ", conjsoma)
}
