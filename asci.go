package ascii

import (
	"strings"
)

func Asci(text, banner string) string {
	content := LoadAscci(banner)
	words := strings.Split(text, "\n")
	if IsEmpty(words) {
		words = words[1:]
	}
	return PrintAscci(words, content)
}
