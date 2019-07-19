package models

import (
	"errors"
	"github.com/lucassamuel/validation"
	"strings"
)

/*
	Tipo de dados utilizados nas funções de
	construção da string de query
*/
type Query struct {
	/*
		Campos selecionados no retorno da consulta

		@var string array
		@example Fields = []string {"id", "subject", "owner"}
	*/
	Fields 			[]string

	/*
		Prefixo incluído na string de seleção dos
		campos da consulta

		@var string
		@example fieldsPrefix = "$select="
	*/
	fieldsPrefix 	string

	/*
		Filtros utilizados na cláusula de condição
		da consulta

		@var string array
		@example Filters = []string {"id=12345", "baseStatus='Em andamento'"}
	*/
	Filters 		[]string

	/*
		Prefixo incluído na string de filtro da
		consulta

		@var string
		@example filtersPrefix = "$filter="
	*/
	filtersPrefix 	string

	/*
		Filtros utilizados na cláusula de expansão
		da consulta

		@var string array
		@example Expand = []string {"owner", "createdBy"}
	*/
	Expand			[]string

	/*
		Prefixo incluído na string de filtro da
		consulta

		@var string
		@example filtersPrefix = "$filter="
	*/
	expandPrefix	string

	/*
		Operadores lógicos utilizados nos filtros
		das consultas

		@var string array
		@example operators = []string {"=", "!=", "=%"}
	*/
	operators 		[]string

	/*
		Operadores relacionais binários que substituem
		os operadores lógicos na construção da string
		da consulta

		@var string array
		@example relationals = []string {"eq", "ne", "like"}
	*/
	relationals 	[]string

	/*
		Resultado final da construção da
		string da consulta

		@var string array
		@example query = []string {fields, filters, queryString}
	*/
	query 			[]string
}

func (q *Query) New(fields []string, filters []string) error {
	if filters == nil {return errors.New("Erro ao construir a query: filtros não podem ser nulos")}

	q.Fields = fields
	q.Filters = filters

	constructErr := q.Construct()

	if constructErr != nil {return constructErr}

	return nil
}

/*
	Construtor da consulta

	Define os prefixos e operadores lógicos e
	relacionais padrões. Além disso, também
	determina o tamanho do array query[]
*/
func (q *Query) Construct() error {
	// Define o prefixo da seleção dos campos e dos filtros
	// TODO IMPLEMENTAR EXPAND
	q.setPrefix("$select=", "$filter=", "&$expand=owner,clients($expand=organization)")

	// Define os operadores lógicos
	operators := []string {"=", "!=", "=%", ">", "<"}

	// Define os operadores relacionais
	relationals := []string {"eq", "ne", "like", "gt", "lt"}

	if len(operators) != len(relationals) {return errors.New("Erro ao construir a query: operadores diferentes dos relacionais")}

	// Atribui os operadores ao objeto
	q.setOperators(operators, relationals)

	// Define o tamanho do array da query
	q.query = []string {"fields", "filters", "queryString"}

	// Chama o construtor da string da query
	q.queryConstructor()

	return nil
}

func (q *Query) GetStringQuery() string {
	return q.query[2]
}

/*
	setPrefix define os prefixos da seleção dos filtros
*/
func (q *Query) setPrefix(fieldPrefix string, filterPrefix string, expandPrefix string) {
	validation.IsEmpty(fieldPrefix, "Erro ao iniciar o construtor: fieldPrefix nulo")
	q.fieldsPrefix = fieldPrefix

	validation.IsEmpty(filterPrefix, "Erro ao iniciar o construtor: filterPrefix nulo")
	q.filtersPrefix = filterPrefix

	q.expandPrefix = expandPrefix
}

func (q *Query) setOperators(operators []string, relationals[]string) {
	q.operators = operators
	q.relationals = relationals
}

func (q *Query) queryConstructor() {
	q.fieldsConstructor()
	q.filtersConstructor()

	q.query[2] = q.query[0] + q.query[1] + q.expandPrefix
}

func (q *Query) fieldsConstructor() {
	if q.Fields[0] != "" {
		q.query[0] = "&" + q.fieldsPrefix

		for i := 0; i < len(q.Fields); i++ {
			if i < (len(q.Fields) - 1) {
				q.query[0] += q.Fields[i] + ","
			} else {
				q.query[0] += q.Fields[i]
			}
		}
	} else {
		q.query[0] = ""
	}
}

func (q *Query) filtersConstructor() {

	q.query[1] = "&"

	if len(q.Filters) > 1 {
		q.query[1] += q.filtersPrefix

		for i := 0; i < len(q.Filters); i++ {
			for j := 0; j < len(q.operators); j++ {
				split := strings.Split(q.Filters[i], q.operators[j])

				if len(split) == 2 {
					q.query[1] += split[0] + "%20" + q.relationals[j] + "%20%27" + strings.ReplaceAll(split[1]," ","%20") + "%27"
				}
			}

			if i < (len(q.Filters) - 1) {
				q.query[1] += "%20and%20"
			}
		}
	} else {
		for i := 0; i < len(q.operators); i++ {
			split := strings.Split(q.Filters[0], q.operators[i])

			if len(split) == 2 {
				if split[0] != "id" || q.operators[i] != "=" {
					q.query[1] += q.filtersPrefix
					break
				}
			}
		}

		for i := 0; i < len(q.operators); i++ {
			split := strings.Split(q.Filters[0], q.operators[i])

			if len(split) == 2 {
				if split[0] == "id" && q.operators[i] == "=" {
					q.query[1] += q.Filters[0]
					break
				} else {
					q.query[1] += split[0] + "%20" + q.relationals[i] + "%20" + strings.ReplaceAll(split[1]," ","%20")
					break
				}
			}
		}
	}
}