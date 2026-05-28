package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type a Sentence")

	Sentence, _ := reader.ReadString('\n')

	Sentence = strings.TrimSpace(Sentence)

	fmt.Println("Length:", len(Sentence))

	fmt.Println("Upper:", strings.ToUpper(Sentence))

	words := strings.Split(Sentence, " ")
	fmt.Println("words:", len(words))

	fmt.Println("Start with Go:", strings.HasPrefix(Sentence, "Go"))

	new := strings.Replace(Sentence, words[0], "Python", 1)
	fmt.Println("Replaced:", new)

	
}