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
}

// URL + TOKEN + SELECT + FILTER
// https://api.movidesk.com/public/v1/tickets?token=MEU_TOKEN&$select=id,subject&$filter=baseStatus eq 'Em pausa' and justification eq 'Versão liberada'

// URL + TOKEN + FILTER
// https://api.movidesk.com/public/v1/tickets?token=MEU_TOKEN&id=22155

// Construtor da string query
func constructor(fields []string, filters []string) string {
	
	var result string
	
	// Define os prefixos da query
	prefix := query {
		// Prefixo dos campos
		fieldsPrefix: "$select=",
		// Prefixo dos filtros
		filtersPrefix: "$filter=",
	}

	// Retorna a string da query
	return result
}

func parseQuery(query query) []string {
	// fields: {"id", "subject", "owner"}
	// filters: {"baseStatus='Em pausa'", "justification='Versão liberada'"}

	// $select=id,subject,owner
	// $filter=baseStatus eq 'Em pausa' and justification eq 'Versão liberada'

	// Variável que armazenará os campos da query
	var fields string
	// Variável que armazenará os filtros da query
	var filters string

	// Validação dos filtros da query, que não podem ser nulos (ao menos um)
	validation.IsNullArray(query.filters, "Erro ao ler a query: filtro nulo")



	for i := 0; i < len(query.fields); i++ {
		if i < (len(query.fields) - 1) {
			fields += query.fields[i] + ","
		} else {
			fields += query.fields[i]
		}
	}
}