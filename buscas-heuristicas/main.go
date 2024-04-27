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
	/*
		fmt.Printf("ALGORITMO GULOSO:\n\n")
		buscaGulosaAteBucharest(romenia, cidadeDeOrigem)
		for _, cidade := range caminhoGuloso {
			fmt.Printf("-> %s ", cidade)
		}

	*/
	fmt.Printf("\n\nALGORITMO A*:\n\n")
	caminhoAEstrela := buscaAEstrelaAteBucharest(romenia, cidadeDeOrigem)
	for _, cidade := range caminhoAEstrela {
		fmt.Printf("-> %s ", cidade)
	}
}

type caminhoPossivel struct {
	caminho []string
	custo   int
}

func buscaAEstrelaAteBucharest(grafo [][]int, cidadeAtual string) []string {
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

			caminhosPossiveis = append(caminhosPossiveis, caminhoPossivel{
				caminho: caminhoAnalisado,
				custo:   novoCusto,
			})

			for _, cp := range caminhosPossiveis {
				fmt.Println(cp.caminho, cp.custo)
			}
			fmt.Println()
			time.Sleep(1 * time.Second)

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
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println()
		fmt.Println()

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

/*
[[A]]

[[A S][A T][A Z]]
   x

[[A S A][A S F][A S O][A S R][A T][A Z]]
                         x

[[A S A][A S F][A S O][A S R P][A S R S][A S R C][A T][A Z]]
           x

[[A S A][A S F S][A S F B][A S O][A S R P][A S R S][A S R C][A T][A Z]]
                                     x

[[A S A][A S F S][A S F B][A S O][A S R P B][A S R P C][A S R P R][A S R S][A S R C][A T][A Z]]
                                      x
*/
