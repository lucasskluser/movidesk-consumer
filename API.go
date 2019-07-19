package movidesk

/*
	Funções responsáveis por chamarem os construtores de consulta e
	requisição, além de obterem todos os dados da resposta ou parte
	dele.

	@author Lucas Samuel
	@version 1.0.0
	@since 17.07.19
*/

import (
	"github.com/lucassamuel/movidesk/models"
	"github.com/pkg/errors"
)

/*
	Estrutura de dados da API do Movidesk
*/
type API struct {
	/*
		Contém o prefixo da URL que será incluído na requisição

		@example https://api.movidesk.com/public/v1/tickets
	*/
	URL      string

	/*
		Contém o token necessário para a autorização das requisições
	*/
	Token    string

	/*
		Contém o corpo da requisição de acordo com a estrutura de
		dados de uma requisição

		Para mais detalhes, leia a documentação da estrutura models.Request
	*/
	Request  models.Request

	/*
		Contém os dados da consulta de acordo com a estrutura de
		dados de uma consulta

		Para mais detalhes, leia a documentação da estrutura models.Query
	*/
	Query    models.Query
}

/*
	NewRequest chama os métodos construtores da query e inicia uma
	nova requisição, passando a URL, o token, a query e o método da
	requisição.

	@param fields []string -> Recebe os campos selecionados da requisição
	@param filters []string -> Recebe os filtros da requisição

	@return error -> Retorna um objeto do tipo error
*/
func (a *API) NewRequest(fields []string, filters []string) error {
	// Chama o construtor da consulta, passando os campos e os filtros como
	// parâmetros, e recebe como retorno um objeto do tipo erro
	errQuery := a.Query.New(fields, filters)

	// Se erro for diferente de nulo, retorna o erro
	if errQuery != nil {return errQuery}

	// Chama o construtor da requisição, passando a URL com o token, a
	// consulta construída e o método da requisição. Recebe como retorno um
	// objeto do tipo erro
	errRequest := a.Request.New(a.URL + "?token=" + a.Token + a.Query.GetStringQuery(), "GET")

	// Se erro for diferente de nulo, retorna o erro
	if errRequest != nil {return errRequest}

	// Se não retornou nenhum outro erro antes, retorna um objeto vazio
	return nil
}

/*
	GetStringRequest retorna a URL da requisição HTTP no formato string
	Essa função é útil para análise da URL que está sendo requisitada

	@return string -> retorna a string da URL completa da requisição
*/
func (a *API) GetStringRequest() string {
	return string(a.Request.URL)
}

/*
	GetAll retorna o vetor completo de tickets da resposta à requisição.

	@return []Ticket 	-> 	retorna o vetor dos tickets contidos na
							resposta da requisição
	@return error
*/
func (a *API) GetAll() ([]models.Ticket, error) {
	var err error

	if len(a.Request.Response.Data) < 1 {
		err = errors.New("Nenhum ticket encontrado")
	}

	return a.Request.Response.Data, err
}

/*
	GetTicket retorna os dados de um determinado ticket de acordo com
	o id do ticket informado.

	Essa função percorre o vetor de tickets e verifica se o parâmetro
	ticketId é igual ao id do ticket naquela posição do vetor.
	Se for, retorna os dados daquele ticket.

	@param ticketId int -> Indica o id do ticket desejado
	@return Ticket -> retorna os dados da estrutura Ticket
*/
func (a *API) GetTicket(ticketId int) (models.Ticket, error) {
	// Declara uma variável do tipo Ticket
	var ticket models.Ticket

	// Percorre o vetor de tickets
	for i := 0; i < len(a.Request.Response.Data); i++ {
		// Se o id informado for igual ao ID do ticket da posição atual no vetor
		if ticketId == a.Request.Response.Data[i].ID {
			// Salva os dados do ticket na variável
			ticket = a.Request.Response.Data[i]
			// Para o loop
			break
		}
	}

	if ticket.Subject == "" {
		return ticket, errors.New("Ticket não encontrado")
	}

	// Retorna os dados do ticket
	return ticket, nil
}
