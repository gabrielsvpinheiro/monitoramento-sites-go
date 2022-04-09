package main

import (
	"fmt"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Iniciando Monitoramento...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Comando Inválido")
	}
}

func exibeIntroducao() {
	nome := "Gustavo"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("0. Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan( &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}


