/*
Geração de URL para requisição GET à API do Movidesk

Este arquivo serve para executar as requisições GET
utilizando o token do Movidesk.

@author Lucas Samuel
@version 0.0.1.07-19
@since 07/2019
*/

package movidesk

/*
Importa os pacotes necessários para a execução da API e validações
*/
import (
	"os"
	validation "github.com/lucassamuel/validation"
	query "github.com/lucassamuel/movidesk/cmd"
	godotenv "github.com/joho/godotenv"
)

// Start carrega as variáveis de ambiente
func New() {
	// Carrega o arquivo .env e, se houver um erro, salva na variável err
	err := godotenv.Load()
	// Se a variável err contiver um erro, retorna um erro
	validation.HasError(err, "Erro ao carregar o .ENV")
}

// Connect serve para gerar a URL de requisição GET
func Connect() string {
	// Obtém o valor da variável URL_MOVIDESK no .env
	url := os.Getenv("URL_MOVIDESK")
	// Valida se a URL é vazia e retorna um erro
	validation.IsEmpty(url, "Erro ao recuperar a variável URL_MOVIDESK")

	// Obtém o valor da variável TOKEN_MOVIDESK no .env
	token := os.Getenv("TOKEN_MOVIDESK")
	// Valida se o token é vazio e retorna um erro
	validation.IsEmpty(token, "Erro ao recuperar a variável URL_TOKEN")

	// Retorna uma string com a url e o token do Movidesk
	// URL_MOVIDESK?token=TOKEN_MOVIDESK
	return url + "?token=" + token
}

// Query retorna a string da URL de requisição GET com os campos e filtros
func Query(fields []string, filters []string) string {
	// Retorna uma string com a url, token, $select e $filter
	// URL_MOVIDESK?token=TOKEN_MOVIDESK&$select=fields&$filters=filters
	return Connect() + query.Constructor(fields, filters)
}