package main

import "fmt"

// Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
// Remova todas as ocorrências desse valor no Slice e imprima o resultado.
func main() {

	slice := []string{"a", "b", "a", "b", "ab", "d", "c", "ac"}
	var valor string
	fmt.Print("Informe um valor: ")
	fmt.Scan(&valor)

	for i := 0; i < len(slice); i++ {
		if valor == slice[i] {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	fmt.Println(slice)
}
