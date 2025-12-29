package main

import (
	"embed"
	"strings"

	spellchecker "github.com/f1monkey/spellchecker/v3"
)

//go:embed francais.txt
var frenchDictFS embed.FS

// frenchWords holds the loaded French words from the embedded file
var frenchWords []string

func init() {
	data, err := frenchDictFS.ReadFile("francais.txt")
	if err != nil {
		panic("failed to read embedded francais.txt: " + err.Error())
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		word := strings.TrimSpace(line)
		if word != "" {
			frenchWords = append(frenchWords, word)
		}
	}
}

// NewFrenchSpellchecker creates and initializes a spellchecker with French dictionary
func NewFrenchSpellchecker() (*spellchecker.Spellchecker, error) {
	// Initialize spellchecker with French alphabet
	sc, err := spellchecker.New(
		"abcdefghijklmnopqrstuvwxyzàâäæçéèêëïîôùûüÿœ", // French alphabet including accented characters
	)
	if err != nil {
		return nil, err
	}

	// Add French words to dictionary
	sc.AddMany(frenchWords)
	return sc, nil
}
