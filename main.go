package main

import (
	"fmt"
)

func main() {

	fmt.Println("Temos as seguintes opções para você: ")
	fmt.Println("1 para soma")
	fmt.Println("2 para subtração")
	fmt.Println("3 para multiplicação")
	fmt.Println("4 para divisão")

	var operador int
	fmt.Print("Qual operador você deseja: ")
	fmt.Scanln(&operador)

	var contador int
	var total float64 = 0
	var aux float64
	if operador == 1 {
		fmt.Print("Ok, você escolheu soma, quantos valores serão somados? ")
		fmt.Scanln(&contador)
		for i := 0; i < contador; i++ {
			fmt.Print("Digite o valor: ")
			fmt.Scanln(&aux)
			total = total + aux
		}
		fmt.Println("O valor total é:", total)

	} else if operador == 2 {
		fmt.Print("Ok, você escolheu subtração, digite o valor maior: ")
		fmt.Scanln(&total)
		fmt.Print("Quantos valores deseja subtrair deste? ")
		fmt.Scanln(&contador)
		for i := 0; i < contador; i++ {
			fmt.Print("Digite o valor: ")
			fmt.Scanln(&aux)
			total = total - aux
		}
		fmt.Println("O valor total é:", total)
	} else if operador == 3 {
		fmt.Print("Ok, você escolheu multiplicação, quantos valores deseja multiplicar? ")
		fmt.Scanln(&contador)
		total = total + 1
		for i := 0; i < contador; i++ {
			fmt.Print("Digite o valor: ")
			fmt.Scanln(&aux)
			total = total * aux
		}
		fmt.Println("O valor total é:", total)
	} else if operador == 4 {
		fmt.Print("Ok, você escolheu divisão, qual valor deseja dividir? ")
		fmt.Scanln(&total)
		fmt.Print("Por qual valor você deseja dividir ", total, "? ")
		fmt.Scanln(&aux)
		total = total / aux
		fmt.Println("O valor total é:", total)
	} else {
		fmt.Println("Valor inválido.")
	}
}
