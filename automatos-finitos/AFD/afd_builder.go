package AFD

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type AFD struct {
	Alfabeto      []string
	Estados       []string
	EstadoInicial string
	EstadosFinais []string
	Transicoes    [][][]string // tratado como uma matriz de adjacências da teoria de grafos
	Testes        Teste
}

type Teste struct {
	Amount int
	Words  []string
}

/* MÉTODOS PUBLICOS */

func New(path string) AFD {
	var afd AFD

	instructions := readEntryFile(path)

	afd = buildAutomata(instructions)

	return afd
}

/* MÉTODOS PRIVADOS */

// readEntryFile recebe arquivo de entrada do autômato, retornando cada linha de instrução
// separada em um "array".
func readEntryFile(path string) []string {
	var entryData []string

	// Abre arquivo em modo leitura.
	file, err := os.Open(path)

	// Lida com erros.
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	// Procura pela quebra de linha (newline) e divide os dados.
	scanner.Split(bufio.ScanLines)

	// Adiciona cada dado em uma posição de um "array".
	for scanner.Scan() {
		entryData = append(entryData, scanner.Text())
	}

	// Fecha arquivo.
	file.Close()

	return entryData
}

// buildAutomata popula a estrutura AFD com as informações trazidas do arquivo de entrada.
func buildAutomata(instructions []string) AFD {
	var afd AFD

	afd.Estados = strings.Fields(instructions[0])  // estados
	afd.Alfabeto = strings.Fields(instructions[1]) // alfabeto

	// Inicializa a matriz de transições
	afd.Transicoes = make([][][]string, len(afd.Estados))
	for i := 0; i < len(afd.Estados); i++ {
		afd.Transicoes[i] = make([][]string, len(afd.Estados))
	}

	// Registra na matriz de transições caminhando por todas linhas que representam
	// transições de um estado.
	for i := 0; i < len(afd.Estados); i++ {
		transicao := strings.Fields(instructions[2+i])

		// Para cada linha, registra as transições (uma para cada símbolo do alfabeto).
		for j := 0; j < len(afd.Alfabeto); j++ {
			stateIndex := stateToIndex(afd.Estados, transicao[j])
			afd.Transicoes[i][stateIndex] = append(afd.Transicoes[i][stateIndex], afd.Alfabeto[j])
		}
	}

	afd.EstadoInicial = strings.Fields(instructions[2+len(afd.Estados)])[0]
	afd.EstadosFinais = strings.Fields(instructions[2+len(afd.Estados)+1])

	afd.Testes.Amount, _ = strconv.Atoi(strings.Fields(instructions[2+len(afd.Estados)+2])[0])
	for i := 0; i < afd.Testes.Amount; i++ {
		word := strings.Fields(instructions[2+len(afd.Estados)+3+i])[0]
		afd.Testes.Words = append(afd.Testes.Words, word)
	}

	return afd
}

// stateToIndex encontra em qual índice se encontra o estado procurado.
func stateToIndex(estados []string, estado string) int {

	for i, v := range estados {
		if v == estado {
			return i
		}
	}

	return -1
}
