package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(("Enter todo"))
	todoTitle, _ := reader.ReadString('\n')
	fmt.Println(todoTitle)
}
