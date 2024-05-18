package ascii

import (
	"bufio"
	"log"
	"os"
)

func LoadAscci(Banner string) (maze []string) {
	f, err := os.Open(Banner + ".txt")
	if err != nil {
		log.Fatalf("error in laoding the banner file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return
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
	return
}
