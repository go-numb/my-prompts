package myprompts

import (
	"fmt"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	assistantName := "ふれんず"
	userName := "おまえ"
	n := "3"

	commands := List()
	for i := 0; i < len(commands); i++ {
		commands[i].Command = strings.ReplaceAll(commands[i].Command, "[[assistant_name]]", assistantName)
		commands[i].Command = strings.ReplaceAll(commands[i].Command, "[[user_name]]", userName)
		commands[i].Command = strings.ReplaceAll(commands[i].Command, "[[n]]", n)

		fmt.Println(commands[i].Title)
		fmt.Println(commands[i].Command)
	}
}

func TestMap(t *testing.T) {
	assistantName := "ふれんず"
	userName := "おまえ"
	n := "3"

	commands := Map()
	for k := range commands {
		cmd := strings.ReplaceAll(commands[k].Command, "[[assistant_name]]", assistantName)
		cmd = strings.ReplaceAll(cmd, "[[user_name]]", userName)
		cmd = strings.ReplaceAll(cmd, "[[n]]", n)
		commands[k] = Prompt{
			Title:   k,
			Command: cmd,
		}

		fmt.Println(k)
		fmt.Println(commands[k].Command)
	}
}
