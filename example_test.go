package cyrslug_test

import (
	"fmt"

	"github.com/aquilax/cyrslug"
)

func ExampleSlug() {
	fmt.Println(cyrslug.Slug("Примерен текст"))
	// Output:
	// primeren-tekst
}
