package ascii

import (
	"strings"
)

func Asci(text, banner string) string {
	// if len(os.Args) != 3 {
	// 	os.Stdout.WriteString("<-----------\nthe usage:\n go run main.go \"text\" fileName\n----------->\n")
	// 	return
	// }
	// if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
	// 	os.Stdout.WriteString("invalid banner\n")
	// 	return
	// }
	// get content from the file in the arguments
	content := LoadAscci(banner)
	// text := os.Args[1]
	words := strings.Split(text, "\\n")
	if IsEmpty(words) {
		words = words[1:]
	}
	return PrintAscci(words, content)
}
