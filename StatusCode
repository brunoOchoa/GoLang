package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	
	exibeIntroducao()

	for {
		exibeMenu()
		
		comando := leComando()

		switch comando {
			case 1:
				iniciaMonitoramento()
			case 2:
				fmt.Println("Exibindo logs...")
			case 0:
				fmt.Println("Saindo do programa")
				os.Exit(0)
			default:
				fmt.Println("Nao conhecee este comando")
				os.Exit(-1)
			}
	}
	

}

	func exibeIntroducao(){
		nome := "Bruno"
		versao := 1.1
		fmt.Println("Olá sr.", nome)
		fmt.Println("A versao do programa é", versao)
	}

	func exibeMenu() {
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibit logs")
		fmt.Println("0 - Sair do programa")
	}

	func leComando() int {
		var comandoLido int
		fmt.Scan(&comandoLido)
		fmt.Println("O comando escolhido foi", comandoLido)

		return comandoLido
	}

	func iniciaMonitoramento() {
		fmt.Println("Monitorando...")
		site := "http://www.brunoochoa.com"
		resp, _ := http.Get(site)
		
		if resp.StatusCode == 200 {
			fmt.Println("SIte:", site, "foi carregado com sucesso")
		} else {
			fmt.Println("Site:", site, "esta com problemas, Statud Code:", resp.StatusCode)
		}
	}
 
