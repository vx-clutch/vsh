package shell

import (
	"fmt"
	"os"
	"strings"
)

func Start() {
	goto shell
shell:
	for {
		command := prompt()
		switch command.name {
		default:
			fmt.Println("Unknown Command")
		case "":
		case " ":
		case "exit":
			fmt.Println("exit")
			break shell
		case "pwd":
			pwd()
		case "ls":
			ls(&command)
		case "cd":
			cd(command.args[0])
		case "clear":
			clear_screen()
		}
	}
}

type command struct {
	name string
	args []string
}

func prompt() command {
	fmt.Print("(vsh) => ")
	var buf string
	fmt.Scanln(&buf)
	i := strings.Split(buf, " ")
	fmt.Println(i, i)
	return command{
		name: i[0],
		args: i[1:],
	}
}

func pwd() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd, inHome := strings.CutPrefix(wd, "/home/")
	if inHome {
		s_pwd := strings.Split(pwd, "")
		current := 0
		for current < len([]string(s_pwd)) {
			char := []string(s_pwd)[current]
			if char != "/" {
				s_pwd = s_pwd[1:]
				current++
			} else {
				s_pwd = s_pwd[2:]
				break
			}
		}
		pwd = strings.Join(s_pwd, "")
	}
	fmt.Println(pwd)
}

func ls(command *command) {
	var dir string
	if len(command.args) == 0 {
		dir = "./"
	} else {
		dir = command.args[0]
	}
	ls, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, e := range ls {
		if e.IsDir() {
			fmt.Print(e.Name() + "/ ")
			continue
		}
		fmt.Println(e.Name() + " ")
	}
}

func cd(path string) {
	var err error
	if path == "" || path == "~" {
		path, err = os.UserHomeDir()
	}
	err = os.Chdir(path)
	if err != nil {
		panic(err)
	}
}

func clear_screen() {
	fmt.Print("\033[H\033[2J")
}
