package main

import (
	"fmt"
)

func analisarNotas(nota1, nota2 float64) (float64, string) {
	media := (nota1 + nota2) / 2

	if media >= 6 {
		return media, "Aprovado"
	} else {
		return media, "Reprovado"
	}
}

func main() {

	nota1 := 5.5
	nota2 := 2.0

	media, situacao := analisarNotas(nota1, nota2)
	fmt.Printf("Média: %.2f\nSituação: %s\n", media, situacao)
}
