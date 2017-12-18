package main

import (
	"fmt"
)

func main() {

	var opcao int

	fmt.Println("Nesta API terão duas opções de ")
	fmt.Println("/mensagem")
	fmt.Println("/usando")

	fmt.Scan(&opcao)

	if opcao == 1 {
	  fmt.Println("Você tem uma nova mensagem no slack, acesse https://hackathonsaude.slack.com/messages")
}else  if opcao == 2 {
	  fmt.Println("Você tem uma nova mensagem no slack, acesse https://hackathonsaude.slack.com/messages")
	  }
	}
