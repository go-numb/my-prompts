package myprompts

import (
	_ "embed"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
)

type Prompt struct {
	Title   string `csv:"title,omitempty"`
	Command string `csv:"command,omitempty"`
}

//go:embed myprompts.csv
var s string

func List() []Prompt {
	commands := make([]Prompt, 0)
	if err := gocsv.Unmarshal(strings.NewReader(s), &commands); err != nil {
		log.Err(err).Msg("csv read error")
		return nil
	}

	return commands
}

func Map() map[string]Prompt {
	commands := make([]Prompt, 0)
	if err := gocsv.Unmarshal(strings.NewReader(s), &commands); err != nil {
		log.Err(err).Msg("csv read error")
		return nil
	}

	var m = make(map[string]Prompt)
	for i := 0; i < len(commands); i++ {
		m[commands[i].Title] = commands[i]
	}

	return m
}
