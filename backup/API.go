package backup

import (
	"errors"
	"github.com/lucassamuel/movidesk/models"
)

type API struct {
	URL      string
	Token    string
	Request  models.Request
	Response models.Response
	Query    models.Query
}

// Select inicializa a construção da query
func (a *API) NewRequest(fd []string, ft []string) error {
	models.New(fd, ft)
	requestErr := models.New(a.URL + "?token=" + a.Token + models.GetQuery(), "GET")

	if requestErr != nil {return requestErr}

	return nil
}

func (a *API) GetStringRequest() string {
	return models.URL
}

/*
	GetAll retorna o vetor completo de tickets da resposta
	à requisição.

	@return []Ticket -> retorna o vetor dos tickets contidos na
						resposta da requisição
*/
func (r *models.Request) GetAll() []models.Ticket {
	return models.Data
}

/*
	GetTicket retorna os dados de um determinado ticket de
	acordo com o id do ticket informado.

	Essa função percorre o vetor de tickets e verifica se
	o parâmetro ticketId é igual ao id do ticket naquela
	posição do vetor. Se for, retorna os dados daquele
	ticket.

	@param ticketId int -> Indica o id do ticket desejado
	@return Ticket -> retorna os dados da estrutura Ticket
*/
func (r *models.Request) GetTicket(ticketId int) (models.Ticket, error) {
	var err error

	// Salva o vetor de tickets
	tickets := models.Data
	// Define uma variável do tipo Ticket
	var ticket models.Ticket

	// Percorre o vetor de tickets
	for i := 0; i < len(tickets); i++ {
		// Se o id informado for igual ao ID do ticket da posição atual no vetor
		if ticketId == models.ID {
			// Salva os dados do ticket na variável
			ticket = tickets[i]
			// Para o loop
			break
		}
	}

	if models.Subject == "" {err = errors.New("Ticket não encontrado")}

	// Retorna os dados do ticket
	return ticket, err
}