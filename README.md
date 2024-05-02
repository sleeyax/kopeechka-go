# kapeechka-go
API wrapper for https://kopeechka.store in go.

## Installation
`go get github.com/sleeyax/kopeechka-go`

## Usage

```go
package main

import (
	"com.github.com/sleeyax/kopeechka-go"
	"log"
)

func main() {
	api := kopeechka.New("your_api_key_here")

	// Check your balance.
	balance, _ := api.GetBalance()
	log.Println("account balance:", balance.Balance)
	
	// Order new email address.
	order, _ := api.OrderMail(kopeechka.OrderMailRequest{Site: "example.com"})
	log.Println("ordered mail:", order.Mail, "order ID:", order.Id)
	
	// Get message.
	message, _ := api.GetMessage(kopeechka.MessageRequest{OrderId: order.Id})
	log.Println("message:", message.Response.Message, "full message:", message.FullMessage)

	// Cancel email
	cancel, _ := api.CancelMail(kopeechka.CancelMailRequest{OrderId: order.Id})
	log.Println("canceled mail status:", cancel.Status)
}
```

See the [complete example](./cmd/main.go) for more details.
