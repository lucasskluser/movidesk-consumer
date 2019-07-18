package cmd

import (
	"strings"
	validation "github.com/lucassamuel/validation"
)

// Tipificação dos dados da query
type query struct {
	// Prefixo dos campos da query
	fieldsPrefix string
	// Campos da query
	fields []string

	// Prefixo dos filtros da query
	filtersPrefix string
	// Filtros da query
	filters []string
	// Operadores da query
	operators []string
	// Operadores relacionais, correspondente a cada operador da query
	relationals []string
}

// URL + TOKEN + SELECT + FILTER
// https://api.movidesk.com/public/v1/tickets?token=MEU_TOKEN&$select=id,subject&$filter=baseStatus eq 'Em pausa' and justification eq 'Versão liberada'

// URL + TOKEN + FILTER
// https://api.movidesk.com/public/v1/tickets?token=MEU_TOKEN&id=22155

// Construtor da string query
func Constructor(fields []string, filters []string) string {
	
	var fieldsResult string
	var filtersResult string
	
	// Define os prefixos da query
	query := query {
		// Prefixo dos campos
		fieldsPrefix: "$select=",
		// Campos selecionados
		fields: fields,
		// Prefixo dos filtros
		filtersPrefix: "$filter=",
		// Filtros indicados
		filters: filters,
		// Operadores da query
		operators: []string {"=", "!=", "=%", ">", "<"},
		// Operadores relacionais
		relationals: []string {"eq", "ne", "like", "gt", "lt"},
	}

	// Salva o resultado da análise dos campos e dos filtros
	fieldsResult, filtersResult = parseQuery(query)

	if fieldsResult != "" {
		// Retorna a string da query
		return "&" + fieldsResult + "&" + filtersResult
	}

	return "&" + filtersResult
}

// Função que analisa a query e retorna as strings dos campos e dos filtros
func parseQuery(query query) (fieldsQuery string, filtersQuery string) {
	// fields: {"id", "subject", "owner"}
	// filters: {"baseStatus='Em pausa'", "justification='Versão liberada'"}

	// $select=id,subject,owner
	// $filter=baseStatus eq 'Em pausa' and justification eq 'Versão liberada'

	// Variável que armazenará o resultado da construção dos campos da query
	var fields string
	// Variável que armazenará o resultado da construção dos filtros da query
	var filters string

	// Validação dos filtros da query, que não podem ser nulos (ao menos um)
	validation.IsNullArray(query.filters, "Erro ao ler a query: filtro nulo")

	// Se o primeiro campo da query for vazio, mas for informado mais de um campo
	if (query.fields[0] == "") && (len(query.fields) > 1) {
		validation.Fatal("Erro ao ler campos da query: fora de posição")
	}

	// Salva o resultado da construção dos campos
	fields += fieldsConstructor(query)
	// Salva o resultado da construção dos filtros
	filters += filtersConstructor(query)

	// Retorna o resultado
	return fields, filters
}

// Função que constrói a string dos campos da query
func fieldsConstructor(query query) string {
	// Variável que armazenará os campos da query
	var fields string

	// Se a primeira posição do array de campos for diferente de ""
	if query.fields[0] != "" {
		// Inicia com o prefixo do select
		fields += query.fieldsPrefix
		// Construtor do select de campos
		for i := 0; i < len(query.fields); i++ {
			
			// Se o índice atual for menor que o número de campos - 1
			if i < (len(query.fields) - 1) {
				// Salva o campo com uma vírgula
				fields += query.fields[i] + ","
			} else {
				// Salva o campo
				fields += query.fields[i]
			}
		}
	}

	// Retorna o resultado
	return fields
}

// Função que constrói a string dos filtros da query
func filtersConstructor(query query) string {
	// Variável que armazenará os filtros da query
	var filters string

	// Se o número de filtros for maior que um
	if len(query.filters) > 1 {
		// Adiciona o prefixo dos filtros
		filters += query.filtersPrefix
	} else {
		for i := 0; i < len(query.operators); i++ {
			// Divide a string daquela posição de array em duas partes com base no operador[i]
			split := strings.Split(query.filters[0], query.operators[i])

			if len(split) == 2 {
				if split[0] != "id" {
					filters += query.filtersPrefix
				}
			}
		}
	}

	// Se a primeira posição do vetor de campos da query for diferente de vazio,
	// ou seja, se houver, no mínimo, um campo informado para o select
	if query.fields[0] != "" {
		// Para cada filtro no array
		for i := 0; i < len(query.filters); i++ {
			// Para cada operador padrão da query
			for j := 0; j < len(query.operators); j++ {
				// Divide a string daquela posição de array em duas partes com base no operador[j]
				split := strings.Split(query.filters[i], query.operators[j])

				if len(split) == 2 {
					// Analisa o operador e retorna a string correspondente
					filters += parseOperator(split, query.operators[j], query.operators, query.relationals)
				}
			}

			// Se o índice atual for menor que o número de filtros - 1
			if i < (len(query.filters) - 1) {
				// Adiciona and ao final da string
				filters += " and "
			}
		}	
	} else {
		if len(query.filters) > 1 {
			validation.Fatal("Erro ao construir os filtros: a cláusula $select é obrigatória quando $filter > 1")
		}

		for i := 0; i < len(query.operators); i++ {
			// Divide a string daquela posição de array em duas partes com base no operador[j]
			split := strings.Split(query.filters[0], query.operators[i])

			if len(split) == 2 {
				if split[0] != "id" {
					validation.Fatal("Erro ao construir os filtros: apenas o id pode ser usado como $filter quando não há $select")
				}

				filters += query.filters[0]
			}
		}
	}

	// Retorna o resultado
	return filters
}

// Função que analisa o operador da string e substitui pelo operador relacional
func parseOperator(str []string, operator string, operators []string, relationals []string) string {
	// Variável que armazenará o resultado
	// Inicia com a primeira parte da string
	result := str[0]

	// Se o número de operadores for diferente do número de relacionais
	if len(operators) != len(relationals) {
		validation.Fatal("Erro ao validar os operadores: número de operadores difere do número de relacionais")
	}

	// Para cada operador da query
	for i := 0; i < len(operators); i++ {
		// Se o operador[i] for igual ao operador da string
		if operators[i] == operator {
			// Substitui pelo operador relacional
			result += " " + relationals[i] + " "
		}
	}

	// Adiciona a segunda parte da string
	result += "'" + str[1] + "'"

	// Retorna o resultado
	return result
}