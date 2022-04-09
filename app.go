package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	fmt.Println("Olá, seja bem vindo ao programa de monitoramento! Digite seu nome: ")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1. Iniciar Monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("0. Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan( &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento...")
	sites := []string{"https://www.google.com.br", "https://www.youtube.com", "https://www.amazon.com"}
	
	for i := 0; i < 5 ; i++{
		fmt.Println("-----------------------------------------------------------")
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
			fmt.Println("-----------------------------------------------------------")
		}
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Monitoramento finalizado.")
	fmt.Println("")	
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

