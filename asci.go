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

func MakeAscii(words, content []string) (res string) {
	for _, word := range words {
		if word == "" {
			res += "\n"
			continue
		}
		for i := 1; i < 9; i++ {
			myLine := ""
			for _, v := range word {
				if v == '\r' {
					continue
				}
				n := int((v - 32) * 9)
				myLine += content[n+i]
			}
			res += myLine + "\n"
		}
	}
	res = RemoveTrailingSpaces(res)
	return
}
