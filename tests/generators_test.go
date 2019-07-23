package tests

import (
	"github.com/lukesamk/movidesk"
	"github.com/lukesamk/movidesk/models"
	"testing"
)

func TestQueryGenerator(t *testing.T) {
	query := new(models.Query)
	errQuery := query.New([]string{"id", "subject", "createdDate"}, []string{"baseStatus=Stopped", "justification=Versão liberada"})

	if errQuery != nil {
		t.Error(errQuery)
	}

	if query.GetStringQuery() != "&$select=id,subject,createdDate&$filter=baseStatus%20eq%20%27Stopped%27%20and%20justification%20eq%20%27Versão%20liberada%27&$expand=owner,clients($expand=organization)" {
		t.Errorf("Erro na geração da query. String gerada: %s", query.GetStringQuery())
	}
}

func TestUrlGenerator(t *testing.T) {
	apiTest := movidesk.New("MY_TOKEN_STRING")

	field := []string{"id", "subject", "createdDate"}
	filter := []string{"baseStatus=Stopped", "justification=Versão liberada"}

	if errRequest := apiTest.NewRequest(field, filter); errRequest != nil {
		t.Error(errRequest)
	}

	if apiTest.GetStringRequest() != "https://api.movidesk.com/public/v1/tickets?token=MY_TOKEN_STRING&$select=id,subject,createdDate&$filter=baseStatus%20eq%20%27Stopped%27%20and%20justification%20eq%20%27Versão%20liberada%27&$expand=owner,clients($expand=organization)" {
		t.Errorf("Erro na geração da URL. URL gerada: %s", apiTest.GetStringRequest())
	}
}