package main

import "fmt"

func main() {
    // Criando um slice de 5 inteiros
    nums := []int{1, 2, 3, 4, 5}

    // Adicionando mais 3 elementos
    nums = append(nums, 6, 7, 8)

    // Imprimindo tudo
    fmt.Println("Slice completo:", nums)
    fmt.Println("Tamanho (len):", len(nums))
    fmt.Println("Capacidade (cap):", cap(nums))
	
}
