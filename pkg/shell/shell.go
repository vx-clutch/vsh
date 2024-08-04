package shell

import (
	"fmt"
	"os/exec"
	"strings"
)

type operation struct {
	command string
	args    []string
}

func (o operation) Exec() error {
	cmd := exec.Command(o.command, o.args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Start() {
	for {
		command := prompt()
		switch strings.TrimSpace(command.command) {
		default:
			// TODO: Detected valid commands
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
