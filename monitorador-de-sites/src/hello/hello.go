package main

import "fmt"
import "os"
import "net/http"
import "time"

const DELAY = 10

func main() {
	exibeInicio()
	for {
		comando := lerComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func exibeInicio() {
	var nome string = "Lucas"
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão ", versao)
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("-- Menu --")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair do programa")
	fmt.Print("Escolha uma opção -> ")
}

func lerComando() int {
	exibeMenu()
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{
		"https://www.alura.com.br/",
		"https://random-status-code.herokuapp.com/",
		"https://www.linkedin.com/",
	}
	for {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(DELAY * time.Second)
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	data := time.Now()
	if resp.StatusCode != 200 {
		fmt.Println(data.Format("[02/January/2006 15:04:05]"), "ERRO! Site:", site, "está com problemas. StatusCode: ", resp.StatusCode)
	} else {
		fmt.Println(data.Format("[02/January/2006 15:04:05]"), "Site: ", site, "foi carregado com sucesso.")
	}
}
