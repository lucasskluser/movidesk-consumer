package main

import (
	"fmt"
	"github.com/joho/godotenv"
<<<<<<< Updated upstream
	"github.com/lucassamuel/movidesk/models"
	"github.com/lucassamuel/validation"
=======
	"github.com/lucassamuel/movidesk"
	"log"
>>>>>>> Stashed changes
	"os"
)

func main() {
<<<<<<< Updated upstream
	err := godotenv.Load(".env")
	validation.HasError(err, "Erro ao carregar o arquivo .env")

	url := os.Getenv("URL_MOVIDESK")
	token := os.Getenv("TOKEN_MOVIDESK")
=======
	errEnv := godotenv.Load("./.env")
	trataErro(errEnv)

	api := movidesk.API{
		URL:   os.Getenv("URL_MOVIDESK"),
		Token: os.Getenv("TOKEN_MOVIDESK"),
	}

	field := []string{"id", "subject"}
	filter := []string{"baseStatus=Stopped", "justification=Versão liberada"}

	errRequest := api.NewRequest(field, filter)
	trataErro(errRequest)
>>>>>>> Stashed changes

	api := models.API{
		URL:   url,
		Token: token,
	}

<<<<<<< Updated upstream
	field := []string{"id", "subject"}
	filter := []string{"id>20066"}

	api.NewRequest(field, filter)
	fmt.Println(api.GetRequestURL())

	api.Request.Run()
	ticket := api.Request.GetAll()
=======
	errResponse := api.Request.Run()
	trataErro(errResponse)

	ticket, errGet := api.GetAll()
	trataErro(errGet)

	for i := 0; i < len(ticket); i++ {
		fmt.Printf("ID: %d, Assunto: %s, Cliente: %s, Organização: %s\n", ticket[i].ID, ticket[i].Subject, ticket[i].Client[0].BusinessName, ticket[i].Client[0].Organization.BusinessName)
	}
}
>>>>>>> Stashed changes

	fmt.Println(ticket)
}
