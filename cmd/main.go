package main

import (
	"kopeechka"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("please provide an API key")
	}

	apiKey := args[1]
	site := "example.com"
	if len(args) > 2 {
		site = args[2]
	}

	api := kopeechka.New(apiKey)

	balance, err := api.GetBalance()
	if err != nil {
		log.Fatalln("failed to log in:", err)
	}

	if balance.Status != kopeechka.Success {
		log.Fatalln("failed to get balance:", balance.Message)
	}

	log.Println("account balance:", balance.Balance)

	log.Println("ordering mail for", site)
	order, err := api.OrderMail(kopeechka.OrderMailRequest{Site: site})
	if err != nil {
		log.Fatalln("failed to order mail:", err)
	}
	log.Println("ordered mail:", order.Mail, order.Id)

	counter := 0

	for {
		if counter == 5 {
			log.Println("didn't receive a message after 5 tries, canceling...")
			break
		}

		message, err := api.GetMessage(kopeechka.MessageRequest{OrderId: order.Id, Full: 1})
		if err != nil {
			log.Fatalln("failed to get message:", err)
		}

		if message.Response.Status == kopeechka.Success {
			log.Println("message:", message.Response.Message, "full message:", message.FullMessage)
			break
		} else {
			log.Println("waiting for message...")
			time.Sleep(time.Second * 5)
		}

		counter++
	}

	log.Println("canceling mail:", order.Mail)

	cancel, err := api.CancelMail(kopeechka.CancelMailRequest{OrderId: order.Id})
	if err != nil {
		log.Fatalln("failed to cancel mail:", err)
	}
	log.Println("canceled mail status:", cancel.Status)
}
