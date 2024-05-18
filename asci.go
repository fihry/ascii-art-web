package ascii

import (
	"strings"
)

func AsciiArt(text, banner string) string {
	content := LoadAscci(banner)
	words := strings.Split(text, "\n")
	if IsEmpty(words) {
		words = words[1:]
	}
	return MakeAscii(words, content)
}
