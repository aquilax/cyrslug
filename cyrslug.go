/*
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

*/
package cyrslug

import (
	"bytes"
	"strings"
)

// Slug functions transliterates the cyrilic text to ascii slug
func Slug(text string) string {
	r := strings.NewReplacer(
		"а", "a",
		"б", "b",
		"в", "v",
		"г", "g",
		"д", "d",
		"е", "e",
		"ж", "j",
		"з", "z",
		"и", "i",
		"й", "y",
		"к", "k",
		"л", "l",
		"м", "m",
		"н", "n",
		"о", "o",
		"п", "p",
		"р", "r",
		"с", "s",
		"т", "t",
		"у", "u",
		"ф", "f",
		"х", "h",
		"ц", "c",
		"ч", "ch",
		"ш", "sh",
		"щ", "sht",
		"ъ", "y",
		"ь", "y",
		"ю", "yu",
		"я", "ya",
		" ", "-",
		"0", "0",
		"1", "1",
		"2", "2",
		"3", "3",
		"4", "4",
		"5", "5",
		"6", "6",
		"7", "7",
		"8", "8",
		"9", "9",
	)
	text = r.Replace(strings.ToLower(text))
	var buffer bytes.Buffer
	var c rune
	for _, c = range text {
		if (c >= 48 && c <= 57) || (c >= 97 && c <= 122) || c == 95 || c == 45 {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}
