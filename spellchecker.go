package main

import (
	spellchecker "github.com/f1monkey/spellchecker/v3"
)

// French words list for dictionary
var frenchWords = []string{
	"bonjour", "merci", "salut", "bonsoir", "oui", "non", "peut", "être", "avoir", "faire",
	"aller", "venir", "dire", "pouvoir", "voir", "savoir", "vouloir", "temps", "année", "jour",
	"homme", "femme", "enfant", "chose", "vie", "monde", "pays", "maison", "ville", "rue",
	"monsieur", "madame", "grand", "petit", "bon", "mauvais", "beau", "nouveau", "vieux", "jeune",
	"français", "français", "anglais", "espagnol", "allemand", "italien", "européen", "américain",
	"travail", "école", "livre", "ordinateur", "téléphone", "voiture", "table", "chaise", "porte", "fenêtre",
	"manger", "boire", "dormir", "parler", "écrire", "lire", "écouter", "regarder", "aimer", "détester",
	"heureux", "triste", "content", "fâché", "surpris", "fatigué", "malade", "sain", "fort", "faible",
	"aujourd", "demain", "hier", "maintenant", "toujours", "jamais", "souvent", "parfois", "rarement",
	"ici", "là", "partout", "nulle", "part", "quelque", "part", "ailleurs", "dedans", "dehors",
	"comment", "pourquoi", "quand", "combien", "lequel", "laquelle", "lesquels", "lesquelles",
	"avec", "sans", "pour", "contre", "chez", "vers", "dans", "sur", "sous", "devant",
	"derrière", "entre", "parmi", "durant", "pendant", "avant", "après", "depuis", "jusqu",
	"café", "thé", "eau", "vin", "pain", "fromage", "viande", "poisson", "légume", "fruit",
	"pomme", "orange", "banane", "tomate", "carotte", "salade", "poulet", "boeuf", "porc",
	"rouge", "bleu", "vert", "jaune", "noir", "blanc", "gris", "rose", "violet", "orange",
	"un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix",
	"vingt", "trente", "quarante", "cinquante", "soixante", "cent", "mille", "million",
	"père", "mère", "fils", "fille", "frère", "soeur", "oncle", "tante", "cousin", "cousine",
	"ami", "amie", "copain", "copine", "voisin", "voisine", "collègue", "patron", "client",
	"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre",
	"novembre", "décembre", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche",
	"matin", "midi", "après", "soir", "nuit", "heure", "minute", "seconde", "semaine", "mois",
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
