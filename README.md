[![wercker status](https://app.wercker.com/status/9fe2ab81e9bf79c172d81730277c8512/m/ "wercker status")](https://app.wercker.com/project/bykey/9fe2ab81e9bf79c172d81730277c8512)
[![GoDoc](https://godoc.org/github.com/aquilax/cyrslug?status.png)](https://godoc.org/github.com/aquilax/cyrslug)

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


