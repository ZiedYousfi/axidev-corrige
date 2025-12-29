package main

import (
	"fmt"
	"log"

	typrio "github.com/ziedyousfi/typr-io-go"
)

func main() {
	fmt.Println("Listening for keyboard events... (Press Space to clear word)")

	// Initialize spellchecker
	sc, err := NewFrenchSpellchecker()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded %d French words into dictionary\n", len(frenchWords))

	// Create window and get text display
	window, text := CreateTypingWindow()

	// Create current word tracker
	cw := NewCurrentWord(text, sc)

	// Start keyboard listener
	listener, err := typrio.NewListener()
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go func() {
		err = listener.Start(cw.Callback)
		if err != nil {
			log.Printf("Listener error: %v", err)
		}
	}()

	window.ShowAndRun()
}