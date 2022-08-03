package main

import (
	"fmt"
	"github.com/byzk-project-deploy/readline"
)

func main() {
	if err := readline.DialRemote("tcp", ":5555"); err != nil {
		fmt.Errorf("An error occurred: %s \n", err.Error())
	}
}
