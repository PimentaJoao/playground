package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const TOTAL_CIDADES = 20

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
	fmt.Printf("Caminho encontrado: ")
	for _, cidade := range caminhoGuloso {
		fmt.Printf("-> %s ", cidade)
	}

	fmt.Printf("\n\nIniciando A*... \n\n")
	time.Sleep(3 * time.Second)

	fmt.Printf("\n\n\n\nALGORITMO A*:\n\n")
	buscaAEstrelaAteBucharest(romenia, cidadeDeOrigem)
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

type caminhoPossivel struct {
	caminho []string
	custo   int
}

func buscaAEstrelaAteBucharest(grafo [][]int, cidadeAtual string) []string {
	// Inicializa árvore de estados.
	caminhosPossiveis := make([]caminhoPossivel, 0, 100)

	solucao := []string{}

	// Inicializa os possíveis caminhos a serem seguidos.
	if len(caminhosPossiveis) == 0 {
		cp := caminhoPossivel{
			// Cidade de origem da busca.
			caminho: []string{cidadeAtual},

			// Peso inicial (nenhuma distância na aresta + distância geométrica até Bucharest).
			custo: 0 + CidadesDistanciaAteBucharest[CidadeParaIndex[cidadeAtual]],
		}

		caminhosPossiveis = append(caminhosPossiveis, cp)
	}

	// Flag que indica se Bucharest foi encontrada ou não.
	encontrou := false

	// Indica que o primeiro caminho analisado é o primeiro dos possíveis.
	// A escolhas das subsequentes análises serão feitas por meio do menor custo de caminho.
	idxCaminhoAnalisado := 0

	for !encontrou {

		caminhoAnalisado := caminhosPossiveis[idxCaminhoAnalisado].caminho

		fmt.Println("caminho analisado agora: ", caminhoAnalisado)
		fmt.Println()
		time.Sleep(1 * time.Second)

		ultimaCidade := caminhoAnalisado[len(caminhoAnalisado)-1]

		if ultimaCidade == "BUCHAREST" {
			encontrou = true
			solucao = caminhoAnalisado
			break
		}

		// Busca no grafo todas as informações de conexão (distância) com a última cidade do caminho analisado.
		conexoesUltimaCidade := grafo[CidadeParaIndex[ultimaCidade]]

		// Controla se o algoritmo deve ou não remover o atual caminho analisado, caso existam novos caminhos
		// encontrados a partir deste.
		deveRemover := false

		for idxCidadeConectada, distancia := range conexoesUltimaCidade {
			// Pula a análise entre cidades que não estão conectadas (distância igual a 0).
			if distancia == 0 {
				continue
			}

			deveRemover = true

			caminhoAnalisado := append(caminhoAnalisado, IndexParaCidade[idxCidadeConectada])
			novoCusto := calculaCusto(grafo, caminhoAnalisado)

			fmt.Println("caminho descoberto: ", caminhoAnalisado)
			fmt.Println("custo calculado:", novoCusto)
			time.Sleep(1 * time.Second)

			caminhosPossiveis = append(caminhosPossiveis, caminhoPossivel{
				caminho: caminhoAnalisado,
				custo:   novoCusto,
			})

		}

		if deveRemover {
			caminhosPossiveis = removeCaminho(caminhosPossiveis, idxCaminhoAnalisado)
		}

		idxCaminhoMenorCusto := 0
		caminhoMenorCusto := 999999999
		for idxCaminho, cam := range caminhosPossiveis {
			if cam.custo < caminhoMenorCusto {
				caminhoMenorCusto = cam.custo
				idxCaminhoMenorCusto = idxCaminho
			}
		}

		idxCaminhoAnalisado = idxCaminhoMenorCusto
		time.Sleep(4 * time.Second)
		fmt.Printf("\n\n\n")
	}

	return solucao
}

func calculaCusto(grafo [][]int, caminho []string) int {
	somaDistanciasCaminho := 0
	for idx := 0; idx < len(caminho); idx++ {
		// Verifica se existe uma cidade depois da atual sendo analisada, para então somar sua distância.
		if len(caminho) == idx+1 {
			break
		}

		indexCidade := CidadeParaIndex[caminho[idx]]
		indexProximaCidade := CidadeParaIndex[caminho[idx+1]]

		somaDistanciasCaminho += grafo[indexCidade][indexProximaCidade]
	}

	distanciaGeometricaUltimaCidade := CidadesDistanciaAteBucharest[CidadeParaIndex[caminho[len(caminho)-1]]]

	return somaDistanciasCaminho + distanciaGeometricaUltimaCidade
}

func removeCaminho(caminhosPossiveis []caminhoPossivel, idx int) []caminhoPossivel {
	return append(caminhosPossiveis[:idx], caminhosPossiveis[idx+1:]...)
}
