package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hi, what's your name? ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s!\n", strings.TrimSuffix(name, "\n"))
}
