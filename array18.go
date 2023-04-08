package main

import "fmt"

func main() {
	var quantidade, i int
	conjPrimos := []int{}
	i = 2
	var primo bool

	fmt.Println("Informe uma quantidade de números primos.")
	fmt.Scanln(&quantidade)

	for quantidade >= len(conjPrimos) {
		primo = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}
		if primo {
			conjPrimos = append(conjPrimos, i)
		}
		i++
	}

	fmt.Println("Os", quantidade, "primeiros números primos são", conjPrimos)
}
