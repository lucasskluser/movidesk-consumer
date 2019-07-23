package models

import (
	"encoding/json"
	"log"
)

/*
	Estrutura de dados de uma resposta de requisição
*/
type Response struct {
	/*
		Corpo da resposta em bytes
	*/
	Body []byte

	/*
		Resposta convertida de bytes para um vetor de tickets
	*/
	Data []Ticket
}

/*
	Read converte o corpo da resposta à requisição em
	um vetor de tickets e retorna todos os dados da resposta.

	Os itens dentro do JSON retornado pela API são convertidos
	em atributos de acordo com a estrutura de dados de um Ticket

	@return Response -> retorna os dados da resposta da requisição
*/
func (self *Response) Read() error {
	// Instancia a estrutura de dados Ticket
	tickets := []Ticket{}
	ticket := Ticket{}

	var jsonErrArray error
	var jsonErrStruct error

	// Converte o JSON de acordo com os dados da estrutura Ticket
	jsonErrArray = json.Unmarshal(self.Body, &tickets)

	if jsonErrArray != nil {
		jsonErrStruct = json.Unmarshal(self.Body, &ticket)
	} else {
		// Salva o vetor de tickets da resposta
		self.Data = tickets
	}

	if (jsonErrArray != nil) && (jsonErrStruct != nil) {
		log.Printf("Não foi possível decodificar a resposta: %v", jsonErrStruct)

		if e, ok := jsonErrStruct.(*json.SyntaxError); ok {
			log.Printf("Erro de sintaxe no byte %d", e.Offset)
		}

		log.Printf("Resposta do Movidesk: %q", self.Body)
	}

	return nil
}

/*
	Estrutura de dados do agrupamento por organização
*/
type GroupByOrganization struct {
	/*
		Contém o nome da organização
	*/
	Nome    string

	/*
		Contém um vetor com todos os tickets da organização
	*/
	Tickets []Ticket
}

/*
	GroupByOrganization agrupa os resultados por organização,
	retornando um vetor de organizações com seus respectivos
	chamados.
*/
func (self *Response) GroupByOrganization() []GroupByOrganization {
	// Cria um vetor para armazenar as organizações
	organizacoes := make([]GroupByOrganization, 0)

	// Para cada ticket presente na resposta da requisição
	for _, ticket := range self.Data {
		// Declara uma variável que define se aquela organização já foi inserida no vetor
		contain := false
		// Para cada organização no vetor
		for key, organizacao := range organizacoes {
			// Se o nome da organização no vetor for igual ao nome da organização no chamado
			if organizacao.Nome == ticket.Client[0].Organization.BusinessName {
				// Adiciona o chamado na posição da organização no vetor
				organizacoes[key].Tickets = append(organizacoes[key].Tickets, ticket)
				// Define contain como verdadeiro
				contain = true
				// Interrompe o loop
				break
			}
		}

		// Se a organização não está no vetor
		if !contain {
			// Define os dados da organização
			organizacao := GroupByOrganization {
				ticket.Client[0].Organization.BusinessName,
				append([]Ticket {ticket}),
			}

			// Adiciona a organização no vetor
			organizacoes = append(organizacoes, organizacao)
		}
	}

	// Retorna o vetor
	return organizacoes
}