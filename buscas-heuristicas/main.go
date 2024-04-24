package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TOTAL_CIDADES = 20

func ImprimeMatriz(m [20][20]int) {
	for i, linha := range m {
		fmt.Printf("%s ", IndexParaCidade[i])
		for _, item := range linha {
			fmt.Printf("%d ", item)
		}
		fmt.Printf("\n")
	}
}

func extraiConexao(conn []string) (int, int, int) {
	city1, ok := CidadeParaIndex[conn[0]]
	if !ok {
		panicMsg := fmt.Sprintf("unknown city: \"%s\"", conn[0])
		panic(panicMsg)

	}
	city2, ok := CidadeParaIndex[conn[2]]
	if !ok {
		panicMsg := fmt.Sprintf("unknown city: \"%s\"", conn[2])
		panic(panicMsg)
	}
	distance, err := strconv.Atoi(conn[1])
	if err != nil {
		panic("unknown distance")
	}
	return city1, city2, distance
}

var CidadeParaIndex = map[string]int{
	"ARAD":           0,
	"BUCHAREST":      1,
	"CRAIOVA":        2,
	"DROBETA":        3,
	"EFORIE":         4,
	"FAGARAS":        5,
	"GIURGIU":        6,
	"HIRSOVA":        7,
	"IASI":           8,
	"LUGOJ":          9,
	"MEHADIA":        10,
	"NEAMT":          11,
	"ORADEA":         12,
	"PITESTI":        13,
	"RIMNICU_VILCEA": 14,
	"SIBIU":          15,
	"TIMISOARA":      16,
	"URZICENI":       17,
	"VASLUI":         18,
	"ZERIND":         19,
}

var IndexParaCidade = inverteMapa(CidadeParaIndex)

func inverteMapa(m map[string]int) map[int]string {
	invertido := make(map[int]string)
	for k, v := range m {
		invertido[v] = k
	}
	return invertido
}

var CidadesDistanciaAteBucharest = []int{
	366, 0, 160, 242, 161, 176, 77, 151, 226, 244, 241, 234, 380, 100, 193, 253, 329, 80, 199, 374,
}

func main() {
	// Instancia uma matriz-grafo da cidade de Romenia.
	romenia := make([][]int, TOTAL_CIDADES)
	for linha := range romenia {
		romenia[linha] = make([]int, TOTAL_CIDADES)
	}

	// Abrindo o arquivo que descreve as conexões.
	file, err := os.Open("./conexoes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Lendo arquivo e montando grafo cidade.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		conexao := strings.Split(line, " ")
		cidade1, cidade2, distancia := extraiConexao(conexao)
		romenia[cidade1][cidade2] = distancia
		romenia[cidade2][cidade1] = distancia
	}

	cidadeDeOrigem := "ARAD"

	fmt.Printf("Caminho da cidade de origem %s até Bucharest...\n\n", cidadeDeOrigem)

	fmt.Printf("ALGORITMO GULOSO:\n\n")
	buscaGulosaAteBucharest(romenia, cidadeDeOrigem)
	for _, cidade := range caminhoGuloso {
		fmt.Printf("-> %s ", cidade)
	}

	fmt.Printf("\n\nALGORITMO A*:\n\n")
	buscaAEstrelaAteBucharest(romenia, cidadeDeOrigem)
	for _, cidade := range caminhoAEstrela {
		fmt.Printf("-> %s ", cidade)
	}
}

var caminhoGuloso = []string{}

func buscaGulosaAteBucharest(grafo [][]int, cidadeAtual string) {
	cidadeDeMenorDistancia := ""
	menorDistancia := 9999999

	fmt.Printf("CAMINHOS ENCONTRADOS: \n\n")

	for i, ligacao := range grafo[CidadeParaIndex[cidadeAtual]] {
		if ligacao > 0 {
			if menorDistancia > CidadesDistanciaAteBucharest[i] {
				menorDistancia = CidadesDistanciaAteBucharest[i]
				cidadeDeMenorDistancia = IndexParaCidade[i]
			}

			// Exibe o estado encontrado nesta camada da árvore de busca, ou seja,
			// encontra as cidades adjacentes à cidade atual da recursividade.
			fmt.Printf("%s\n%d milhas até Bucharest\n\n", IndexParaCidade[i], CidadesDistanciaAteBucharest[i])
		}
	}

	fmt.Println("**********************************")
	fmt.Println("RESULTADOS: ")
	fmt.Println("cidade de menor distancia: ", cidadeDeMenorDistancia)
	fmt.Println(menorDistancia, "milhas")
	fmt.Println("**********************************")
	fmt.Println("")

	caminhoGuloso = append(caminhoGuloso, cidadeAtual)

	if menorDistancia > 0 {
		buscaGulosaAteBucharest(grafo, cidadeDeMenorDistancia)
	} else if menorDistancia == 0 {
		caminhoGuloso = append(caminhoGuloso, cidadeDeMenorDistancia)
	}
}

var caminhoAEstrela = []string{}

func buscaAEstrelaAteBucharest(grafo [][]int, cidadeAtual string) {
	if cidadeAtual == "BUCHAREST" {
		caminhoAEstrela = append(caminhoAEstrela, "BUCHAREST")
		return
	}

	cidadeDeMenorDistancia := ""
	menorDistancia := 9999999

	fmt.Printf("CAMINHOS ENCONTRADOS: \n\n")

	for i, ligacao := range grafo[CidadeParaIndex[cidadeAtual]] {
		if ligacao > 0 {
			if menorDistancia > CidadesDistanciaAteBucharest[i]+ligacao {
				menorDistancia = CidadesDistanciaAteBucharest[i] + ligacao
				cidadeDeMenorDistancia = IndexParaCidade[i]
			}

			// Exibe o estado encontrado nesta camada da árvore de busca, ou seja,
			// encontra as cidades adjacentes à cidade atual da recursividade.
			fmt.Printf("%s\n%d milhas até Bucharest\n\n", IndexParaCidade[i], CidadesDistanciaAteBucharest[i]+ligacao)
		}
	}

	fmt.Println("**********************************")
	fmt.Println("RESULTADOS: ")
	fmt.Println("cidade de menor distancia composta: ", cidadeDeMenorDistancia)
	fmt.Println(menorDistancia, "milhas")
	fmt.Println("**********************************")
	fmt.Println("")

	caminhoAEstrela = append(caminhoAEstrela, cidadeAtual)
	buscaAEstrelaAteBucharest(grafo, cidadeDeMenorDistancia)
}
