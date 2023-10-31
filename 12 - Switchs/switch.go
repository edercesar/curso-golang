package main

import "fmt"

// 1 modo de usar
func DiadaSemana(dia int) string {
	switch dia {
	case 1:
		return ("Domingo")
	default:
		return ("Dia invalido")
	}

}

// 2 modo de usar
func DiadaSemana2(dia2 int) string {
	switch {
	case dia2 == 1:
		return "Hoje e domingo"
	default:
		return "escolheu numero errado"
	}

}

// 3 modo de usar

func DiadaSemana3(dia3 int) string {
	var DiadaSemana string
	switch {
	case dia3 == 1:
		DiadaSemana = "Domingo"
	case dia3 == 2:
		DiadaSemana = "Segunda-Feira"
	case dia3 == 3:
		DiadaSemana = "Terca-Feira"
	case dia3 == 4:
		DiadaSemana = "Quarta-Feira"
		fallthrough
	case dia3 == 5:
		DiadaSemana = "Quinta-feira"
	default:
		DiadaSemana = "Numero INvalido"
	}

	return DiadaSemana
}

func main() {
	fmt.Println("Aula de Switch Jeito 1")
	dia := DiadaSemana(1)
	fmt.Println(dia)
	fmt.Println("Aula de Switch Jeito 2")
	dia2 := DiadaSemana2(2)
	fmt.Println(dia2)
	fmt.Println("Aula de Switch Jeito 3")
	dia3 := DiadaSemana3(4)
	fmt.Println(dia3)

}
