Package cyrslug generates transliterated ascii slug from cyrilic titles

Exaple:

	package main

	import (
		"github.com/aquilax/cyrslug"
	)

	func main() {
		text := "Примерен текст"
		print(cyrslug.Slug(text)) // Outputs: "primeren-tekst"
	}


