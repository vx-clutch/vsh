package shell

import (
	"fmt"
	"strings"
)

type operation struct {
	command string
	args    []string
}

func Start() {
	for {
		comand := prompt()
		switch comand.command {
		default:
			fmt.Print("vsh: Unknown command")
		case "", " ":
			continue
		case "exit":
			fmt.Println("exit")
			return

		}
		fmt.Println()
	}
}

func prompt() operation {
	fmt.Print("(vsh) => ")
	var input_buf string
	fmt.Scanln(&input_buf)
	buf := strings.Split(input_buf, " ")
	return operation{
		command: buf[0],
		args:    buf[1:],
	}
}
