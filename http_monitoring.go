package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 2
const delay = 2

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa.")
		// } else {
		// 	fmt.Println("Comando não identificado, saindo do programa!")
		// }

		// variavel comando recebe o conteudo int presente na funcao leComando
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa.")
			// 0 informa ao OS que saiu com sucesso
			os.Exit(0)
		default:
			fmt.Println("Comando não identificado, saindo do programa!")
			// -1 indica algo inesperado no codigo, nesse caso, um numero inesperado do menu
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	var nome string
	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Programa na versao:", versao)
	fmt.Println("")
	fmt.Println("Bem vindo,", nome, ", escolha sua opção no Menu abaixo:")
	fmt.Println("")

}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento.")
	fmt.Println("2- Exibir logs de monitoramento.")
	fmt.Println("0- Sair do programa.")
}

func leComando() int {
	var comandoLido int
	//posso usar o scanf citando o tipo ou somente o scan e deixar o go inferir o tipo
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	leSitesDoArquivo()
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(i, site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(i int, site string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site", i, ":", site, "abriu com sucesso.")
		registraLog(site, true)
	} else {
		fmt.Println("Site", i, ":", site, "não abriu, o código de erro é:", res.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	// se existir algum valor presente dentro da err, vai printar aqui
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	// le e printa o arquivo linha por linha ate bater no erro EOF (end of file)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro ao gerar o log:", err)
	}
	arquivo.WriteString(time.Now().Format("2006/01/02 | 15:04:05") + " | " + site + " | online: " + strconv.FormatBool(status)) //+ "\n")
	fmt.Fprintln(arquivo)
	arquivo.Close()
}

func imprimeLogs() {
	// ioutil nao precisa fechar o arquivo, le tudo e devolve so o array de bytes, nivel mais alto que OS que eh a nivel de sistema operacional
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao imprimir os logs: ", err)
	}
	// convertendo os bytes que a ioutil leu pra string
	fmt.Println(string(arquivo))
}
