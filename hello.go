package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando!")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Vinicius"
	versao := 1.1
	fmt.Println("Ola sr. ", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

}