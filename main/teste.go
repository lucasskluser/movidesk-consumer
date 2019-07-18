package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lucassamuel/movidesk/models"
	"github.com/lucassamuel/validation"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	validation.HasError(err, "Erro ao carregar o arquivo .env")

	url := os.Getenv("URL_MOVIDESK")
	token := os.Getenv("TOKEN_MOVIDESK")

	api := models.API{
		URL:   url,
		Token: token,
	}

	field := []string{""}
	filter := []string{"id=20066"}

	api.NewRequest(field, filter)
	fmt.Println(api.GetRequestURL())

	api.Request.Run()
	ticket := api.Request.GetAll()

	fmt.Println(ticket)
}
