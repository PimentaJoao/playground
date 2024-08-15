package AFD

type TestHandler struct {
	Word           string
	States         []string
	AcceptedStatus bool
}

// Process recebe uma um AFD, retornando um "array" de estados percorridos e
// se suas palavras foram aceita ou não.
func Process(afd AFD) []TestHandler {
	var tests []TestHandler

	// Para cada palavra, encontra se ela foi aceita e por quais estados passou.
	for i := 0; i < afd.Testes.Amount; i++ {

		// Inicializa "array" de estados com o estado inicial.
		var states []string = nil
		states = append(states, afd.EstadoInicial)

		// Percorre todos os símbolos de uma palavra, registrando os estados percorridos.
		for j := 0; j < len(afd.Testes.Words[i]); j++ {
			currentState := states[len(states)-1]
			nextState := discoverNextState(afd, currentState, string(afd.Testes.Words[i][j]))
			states = append(states, nextState)
		}

		// Verifica se a palavra foi aceita ou não
		accepted := wordAccepted(afd, states)

		// Monta um estrutura de resposta.
		test := TestHandler{
			Word:           afd.Testes.Words[i],
			States:         states,
			AcceptedStatus: accepted,
		}

		// Guarda a estrutura de resposta para aquela palavra.
		tests = append(tests, test)
	}

	return tests
}

func discoverNextState(afd AFD, currentState string, symbol string) string {
	// Encontra o índice do estado onde a palavra está agora
	stateIndex := stateToIndex(afd.Estados, currentState)

	// Dentro desse índice, procura e retorna para qual estado o símbolo atual da
	// palavra "aponta".
	for i := 0; i < len(afd.Estados); i++ {
		for _, s := range afd.Transicoes[stateIndex][i] {
			if s == symbol {
				return afd.Estados[i]
			}
		}
	}

	return "NOT_FOUND"
}

func wordAccepted(afd AFD, states []string) bool {
	lastState := states[len(states)-1]

	// Verifica se o último estado alcançado fazia parte do conjunto dos estados finais.
	for _, e := range afd.EstadosFinais {
		if e == lastState {
			return true
		}
	}

	return false
}
