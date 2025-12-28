package main

import (
	"fmt"
	"log"

	typrio "github.com/ziedyousfi/typr-io-go"
)

func main() {
	fmt.Println("Hello, World!")

	sender, err := typrio.NewSender()
	if err != nil {
		log.Fatal(err)
	}
	defer sender.Close()

	// Type some text
	sender.TypeText("Hello, World!")
	// Press a key combination
	sender.Combo(typrio.ModCtrl, typrio.StringToKey("S"))
}
