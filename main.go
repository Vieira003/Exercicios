package main

import "fmt"

func main() {

    nums := []int{1, 2, 3, 4, 5}

    nums = append(nums, 6, 7, 8)

    fmt.Println("Slice completo:", nums)
    fmt.Println("Tamanho len:", len(nums))
    fmt.Println("Capacidade cap:", cap(nums))
	
    nomes := []string{"Alice", "Bruno", "Carlos", "Daniela", "Eduardo"}

    fmt.Println("Dois primeiros:", nomes[:2])

    fmt.Println("Dois últimos:", nomes[len(nomes)-2:])

    fmt.Println("Nome do meio:", nomes[2:3])
}
