package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter one or more words: ")
	words, _ := reader.ReadString('\n')
	trimmed := strings.TrimSuffix(words, "\n")
	fmt.Printf("The text '%s' has %d characters (not including spaces)\n",
		trimmed, len(strings.Replace(trimmed, " ", "", -1)))
}
