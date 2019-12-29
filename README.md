[![GoDoc](https://godoc.org/github.com/aquilax/cyrslug?status.png)](https://godoc.org/github.com/aquilax/cyrslug)

Package cyrslug generates transliterated ascii slug from cyrillic titles

Example:

```go
package main

import (
	"github.com/aquilax/cyrslug"
)

func main() {
	text := "Примерен текст"
	print(cyrslug.Slug(text)) // Outputs: "primeren-tekst"
}
```
