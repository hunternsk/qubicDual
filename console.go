package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func promptUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt, " -> ")
	variable, _ := reader.ReadString('\n')
	variable = strings.Replace(variable, "\r", "", -1)
	variable = strings.Replace(variable, "\n", "", -1)
	if len(variable) == 0 {
		fmt.Println("Must not be empty")
		os.Exit(1)
	}
	return variable
}
