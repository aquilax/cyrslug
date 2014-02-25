package cyrslug

import (
	"testing"
)

func TestSlug(t *testing.T) {
	cases := map[string]string{
		"АБВ":                              "abv",
		"Сайт на деня":                     "sayt-na-denya",
		"Какво да купя?":                   "kakvo-da-kupya",
		"Търся MP3 плейър за 160/170 лева": "tyrsya-mp3-pleyyr-za-160170-leva",
	}
	for text, expected := range cases {
		slug := Slug(text)
		if expected != slug {
			t.Error("Expected "+expected+", got", slug)
		}
	}
}
