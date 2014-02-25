package cyrslug

import (
	"bytes"
	"strings"
)

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
	var o rune
	for _, c := range text {
		o = c
		if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 95 || o == 45 {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}
