package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/lukesamk/movidesk"
)

func main() {
	api := movidesk.New("ec673f86-3d7a-4045-afc0-4b453e148678")

	field := []string{"id", "subject", "createdDate"}
	filter := []string{"baseStatus=Stopped", "justification=Versão liberada"}

	errRequest := api.NewRequest(field, filter)
	trataErro(errRequest)

	fmt.Println(api.GetStringRequest())

	errResponse := api.Request.Run()
	trataErro(errResponse)

	groupBy := api.Request.Response.GroupByOrganization()

	for _, organizacao := range groupBy {
		fmt.Printf("%s:\n", organizacao.Nome)
		for _, ticket := range organizacao.Tickets {
			createdDate := strings.Split(ticket.CreatedDate, "T")
			fmt.Printf(
				"%d [%s] - %s (%s) / %s\n",
				ticket.ID,
				createdDate[0],
				ticket.Subject,
				ticket.Client[0].BusinessName,
				ticket.Owner.BusinessName,
			)
		}
		fmt.Print("\n")
	}
}

func trataErro(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
