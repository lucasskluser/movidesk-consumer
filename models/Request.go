package models

/*
	Funções responsáveis por executarem as requisições à API
	do Movidesk e retornarem a resposta convertida para o tipo
	de dados incluídos em um chamado

	@author Lucas Samuel
	@version 1.0.0
	@since 17.07.19
*/

import (
	"encoding/json"
	"github.com/lucassamuel/validation"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
	Estrutura de dados contidos em uma requisição
*/
type Request struct {
	/*
		Contém a URL onde será efetuada a requisição. O token,
		os campos e os filtros são incluídos aqui

		@example https://api.movidesk.com/public/v1/tickets?token=TOKEN&$select=id,subject&$filters=createdDate eq 2019-01-01
	*/
	URL string

	/*
		Indica o método que será utilizado na requisição

		@example GET
	*/
	// TODO: Implementar funções para o método POST
	Method string

	/*
		Contém a resposta da requisição de acordo com a
		estrutura de dados de uma resposta de requisição
	*/
	Response Response
}

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
	New inicia uma nova requisição à API

	Essa função valida se a URL e o método informados são vazios e retorna
	um erro se forem. Se não forem vazios, define a URL e o método da
	requisição de acordo com os parâmetros da função.

	@param url string -> Recebe a URL que será utilizada para a requisição
	@param method string -> Indica o tipo de requisição que será utilizado
*/
<<<<<<< Updated upstream
func (r *Request) New(url string, method string) {
	// Verifica se a URL é vazia e retorna um erro
	validation.IsEmpty(url, "Erro ao definir a URL do Request: url vazia")
	// Verifica se o método é vazio e retorna um erro
	validation.IsEmpty(method, "Erro ao definir o método do Request: método vazio")
=======
func (r *Request) New(url string, method string) error {
	// Variável para retornar erro
	var err error

	// Se a URL é vazia, retorna um erro
	if url == "" {err = errors.New("Erro ao criar uma requisição: a URL não pode ser vazia"); return err}

	// Se o método é vazio, retorna um erro
	if method == "" {err = errors.New("Erro ao criar uma requisição: o método não pode ser vazio"); return err}
>>>>>>> Stashed changes

	// Define a URL da requisição
	r.URL = url
	// Define o método da requisição
	r.Method = method
}

/*
	Run executa a requisição à API

	O método e a URL da requisição são obtidos dos dados que foram
	informados através do início da requisição (função New).

	Se houver um erro durante a execução da requisição, ele é retornado
	como erro fatal, encerrando a aplicação. O corpo da resposta à
	requisição é salvo e convertido para um vetor de tickets.
*/
func (r *Request) Run() {
	// Define o timeout da requisição em 5 segundos
<<<<<<< Updated upstream
	client := http.Client{Timeout: time.Second * 5}
=======
	client := http.Client{Timeout: time.Second * 10}

>>>>>>> Stashed changes
	// Define a URL e o método da requisição de acordo com o que foi passado pela função New
	request, clientErr := http.NewRequest(r.Method, r.URL, nil)
	//log.Fatal(request)
	// Valida se houve algum erro durante a definição da requisição
	validation.HasError(clientErr, "Erro ao definir a requisição HTTP")
	// Define o cabeçalho da requisição
	request.Header.Set("User-Agent", "movidesk-api")
	// Executa a requisição
	response, responseErr := client.Do(request)
	// Valida se houve algum erro durante a execução da requisição
	validation.HasError(responseErr, "Erro ao executar a requisição HTTP")

	// Recebe o corpo da resposta à requisição
	body, readErr := ioutil.ReadAll(response.Body)
	// Valida se houve algum erro durante a leitura da resposta da requisição
	validation.HasError(readErr, "Erro ao ler a resposta da requisição HTTP")
	// Se o corpo da resposta for vazio
	if body == nil {
		// Retorna um erro
		log.Fatal(errors.Errorf("Erro ao ler o corpo da resposta: %s"), body)
	}

	// Salva o corpo da resposta
	r.Response.Body = body
	// Converte a resposta em um vetor de tickets
	r.ReadResponse()
}

/*
	ReadResponse converte o corpo da resposta à requisição em
	um vetor de tickets e retorna todos os dados da resposta.

	Os itens dentro do JSON retornado pela API são convertidos
	em atributos de acordo com a estrutura de dados de um Ticket

	@return Response -> retorna os dados da resposta da requisição
*/
func (r *Request) ReadResponse() Response {
	// Instancia a estrutura de dados Ticket
<<<<<<< Updated upstream
	ticket := []Ticket{}
	// Converte o JSON de acordo com os dados da estrutura Ticket
	jsonErr := json.Unmarshal(r.Response.Body, &ticket)
	// Verifica se houve algum erro durante a conversão
	if jsonErr != nil {
		log.Printf("Erro ao decodificar a resposta: %v", jsonErr)
		if e, ok := jsonErr.(*json.SyntaxError); ok {
			log.Printf("Erro de sintaxe no byte %d", e.Offset)
		}
		log.Printf("Movidesk response: %q", r.Response.Body)
=======
	tickets := []Ticket{}
	ticket := Ticket{}

	// Converte o JSON de acordo com os dados da estrutura Ticket
	jsonErr := json.Unmarshal(r.Response.Body, &tickets)

	// Contorno temporário
	if jsonErr != nil {
		log.Printf("Não foi possível decodificar a resposta na estrutura []Ticket: %v\nTentando decodificar na estrutura Ticket", jsonErr)
		jsonErr = json.Unmarshal(r.Response.Body, &ticket)

		// Verifica se houve algum erro durante a conversão
		if jsonErr != nil {
			log.Printf("Erro ao decodificar a resposta: %v", jsonErr)

			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("Erro de sintaxe no byte %d", e.Offset)
			}

			log.Printf("Movidesk response: %q", r.Response.Body)
		}

		// Salva o vetor de tickets da resposta
		r.Response.Data = []Ticket {ticket}

		return r.Response
>>>>>>> Stashed changes
	}
	// Salva o vetor de tickets da resposta
<<<<<<< Updated upstream
	r.Response.Data = ticket
	// Retorna a resposta completa
=======
	r.Response.Data = tickets

>>>>>>> Stashed changes
	return r.Response
}

/*
	GetAll retorna o vetor completo de tickets da resposta
	à requisição.

	@return []Ticket -> retorna o vetor dos tickets contidos na
						resposta da requisição
*/
func (r *Request) GetAll() []Ticket{
	return r.Response.Data
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
func (r *Request) GetTicket(ticketId int) Ticket {
	// Salva o vetor de tickets
	tickets := r.Response.Data
	// Define uma variável do tipo Ticket
	var ticket Ticket

	// Percorre o vetor de tickets
	for i := 0; i < len(tickets); i++ {
		// Se o id informado for igual ao ID do ticket da posição atual no vetor
		if ticketId == tickets[i].ID {
			// Salva os dados do ticket na variável
			ticket = tickets[i]
			// Para o loop
			break
		}
	}
	// Retorna os dados do ticket
	return ticket
}