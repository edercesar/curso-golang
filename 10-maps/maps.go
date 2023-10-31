package main

import "fmt"

func main() {
	fmt.Println("Aula de MAPS")
	// criando um usuario usando um mapa
	usuario1 := map[string]string{
		"nome":      "Eder",
		"sobrenome": "Luz",
	}

	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"], usuario1["sobrenome"])

	// criando um mapa que tem como valor um outro mapa

	usuario2 := map[string]map[string]string{
		"pessoa1": {
			"nome":      "Rafael",
			"sobrenome": "Luz",
		},
		"pessoa2": {
			"nome":      "Joao",
			"sobrenome": "Eduardo",
		},
		"pessoa3": {
			"nome":      "Mariana",
			"sobrenome": "Salles",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "pessoa3")
	fmt.Println(usuario2)
}
