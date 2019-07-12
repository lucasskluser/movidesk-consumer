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
	"strings"
	validation "github.com/lucassamuel/validation"
)

/*
	Connect serve para gerar a URL de requisição GET
	@return string
*/
func connect() string {
	// Obtém o valor da variável URL_MOVIDESK no .env
	url := os.Getenv("URL_MOVIDESK")
	validation.IsEmpty(url, "Erro ao recuperar a variável URL_MOVIDESK")

	// Obtém o valor da variável TOKEN_MOVIDESK no .env
	token := os.Getenv("TOKEN_MOVIDESK")
	validation.IsEmpty(token, "Erro ao recuperar a variável URL_TOKEN")

	// URL_MOVIDESK?token=TOKEN_MOVIDESK
	return url + "?token=" + token
}