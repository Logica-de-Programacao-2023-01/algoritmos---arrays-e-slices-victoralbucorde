package main

import "fmt"

// Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
// Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.
func main() {

	slice := []int{1, 2, 3, 4, 5}
	var numero int
	confirmador := false
	fmt.Print("Digite um número para ser adicionado na lista: ")
	fmt.Scan(&numero)

	for i := 0; i < len(slice); i++ {
		if numero == slice[i] {
			confirmador = true
		}
	}
	if confirmador {
		fmt.Println("O número já está presente na lista")
		fmt.Println(slice)
	} else {
		slice = append(slice, numero)
		fmt.Println(slice)
	}

}
