package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	trataErro(err)

	api := API {
		URL:   os.Getenv("URL_MOVIDESK"),
		Token: os.Getenv("TOKEN_MOVIDESK"),
	}

	field := []string{""}
	filter := []string{""}

	err = api.NewRequest(field, filter)
	trataErro(err)

	fmt.Println(api.GetStringRequest())

	err = api.Request.Run()
	trataErro(err)

	ticket := api.GetTicket(20066)

	fmt.Println(ticket.Subject)
}

func trataErro(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
