package main

import (
	"fmt"
	"io/ioutil"

	"github.com/PimentaJoao/DFA-go/AFD"
)

func main() {

	// Lê o diretório "/exemples".
	files, _ := ioutil.ReadDir("./examples")

	// Para cada arquivo encontrado, cria seu AFD e analisa suas palavras teste.
	for _, file := range files {
		fmt.Print("\nARQUIVO: ", file.Name(), "\n\n\n")

		// Constrói o caminho para o arquivo encontrado.
		path := "./examples/" + file.Name()

		// Constrói o AFD.
		afd := AFD.New(path)

		// Testa os exemplos que vieram com o AFD.
		wordTest := AFD.Process(afd)

		// Imprime resultado dos testes.
		for _, test := range wordTest {
			fmt.Print("Palavra: ", test.Word, "\n")
			fmt.Println("Estados percorridos: ")
			for _, w := range test.States {
				fmt.Print(" -> ", w)
			}
			fmt.Print("\nPalavra foi aceita?: ", test.AcceptedStatus, "\n\n")
		}
	}

}
