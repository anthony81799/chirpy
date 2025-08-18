package main

import "strings"

func cleanBody(body string) string {
	words := strings.Fields(body)
	for i, word := range words {
		if strings.ToLower(word) == "kerfuffle" || strings.ToLower(word) == "sharbert" || strings.ToLower(word) == "fornax" {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
