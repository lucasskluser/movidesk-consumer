package models

import (
	"errors"
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
	fieldsPrefix string

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
	filtersPrefix string

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

func (self *Query) New(fields []string, filters []string) error {
	if filters == nil {return errors.New("Erro ao construir a query: filtros não podem ser nulos")}

	self.Fields = fields
	self.Filters = filters

	constructErr := self.Construct()

	if constructErr != nil {return constructErr}

	return nil
}

/*
	Construtor da consulta

	Define os prefixos e operadores lógicos e
	relacionais padrões. Além disso, também
	determina o tamanho do array query[]
*/
func (self *Query) Construct() error {
	// Define o prefixo da seleção dos campos e dos filtros
	// TODO IMPLEMENTAR EXPAND
	errPrefix := self.setPrefix("$select=", "$filter=", "&$expand=owner,clients($expand=organization)")
	if errPrefix != nil {return errPrefix}

	// Define os operadores lógicos
	operators := []string {"=", "!=", "=%", ">", "<"}

	// Define os operadores relacionais
	relationals := []string {"eq", "ne", "like", "gt", "lt"}

	if len(operators) != len(relationals) {return errors.New("Erro ao construir a query: operadores diferentes dos relacionais")}

	// Atribui os operadores ao objeto
	self.setOperators(operators, relationals)

	// Define o tamanho do array da query
	self.query = []string {"fields", "filters", "queryString"}

	// Chama o construtor da string da query
	self.queryConstructor()

	return nil
}

func (self *Query) GetStringQuery() string {
	return self.query[2]
}

func (self *Query) GetQuery(index int) string {
	return self.query[index]
}

/*
	setPrefix define os prefixos da seleção dos filtros
*/
func (self *Query) setPrefix(fieldPrefix string, filterPrefix string, expandPrefix string) error {
	if fieldPrefix == "" {return errors.New("Erro ao iniciar o construtor: fieldPrefix nulo")}
	self.fieldsPrefix = fieldPrefix

	if filterPrefix == "" {return errors.New("Erro ao iniciar o construtor: filterPrefix nulo")}
	self.filtersPrefix = filterPrefix

	self.expandPrefix = expandPrefix

	return nil
}

func (self *Query) setOperators(operators []string, relationals[]string) {
	self.operators = operators
	self.relationals = relationals
}

func (self *Query) queryConstructor() {
	self.fieldsConstructor()
	self.filtersConstructor()

	self.query[2] = self.query[0] + self.query[1] + self.expandPrefix
}

func (self *Query) fieldsConstructor() {
	if self.Fields[0] != "" {
		self.query[0] = "&" + self.fieldsPrefix

		for i := 0; i < len(self.Fields); i++ {
			if i < (len(self.Fields) - 1) {
				self.query[0] += self.Fields[i] + ","
			} else {
				self.query[0] += self.Fields[i]
			}
		}
	} else {
		self.query[0] = ""
	}
}

func (self *Query) filtersConstructor() {

	self.query[1] = "&"

	if len(self.Filters) > 1 {
		self.query[1] += self.filtersPrefix

		for i := 0; i < len(self.Filters); i++ {
			for j := 0; j < len(self.operators); j++ {
				split := strings.Split(self.Filters[i], self.operators[j])

				if len(split) == 2 {
					self.query[1] += split[0] + "%20" + self.relationals[j] + "%20%27" + strings.ReplaceAll(split[1]," ","%20") + "%27"
				}
			}

			if i < (len(self.Filters) - 1) {
				self.query[1] += "%20and%20"
			}
		}
	} else {
		for i := 0; i < len(self.operators); i++ {
			split := strings.Split(self.Filters[0], self.operators[i])

			if len(split) == 2 {
				if split[0] != "id" || self.operators[i] != "=" {
					self.query[1] += self.filtersPrefix
					break
				}
			}
		}

		for i := 0; i < len(self.operators); i++ {
			split := strings.Split(self.Filters[0], self.operators[i])

			if len(split) == 2 {
				if split[0] == "id" && self.operators[i] == "=" {
					self.query[1] += self.Filters[0]
					break
				} else {
					self.query[1] += split[0] + "%20" + self.relationals[i] + "%20" + strings.ReplaceAll(split[1]," ","%20")
					break
				}
			}
		}
	}
}