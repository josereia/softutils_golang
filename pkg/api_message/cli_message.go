package messages

import (
	"log"
	"os"
)

/*
[err] title
description
solve
doc
stack
*/

type CliConfigFunc func(*CliMessage)

type CliMessage struct {
	Title   string
	Doc     string
	Solve   string
	Message Message
}

func NewCliMessage(title string, configs ...MessageFunc) CliMessage {
	msgConfig := Message{}
	for _, fn := range configs {
		fn(&msgConfig)
	}

	return CliMessage{
		Title:   title,
		Message: msgConfig,
	}
}

func (cfg *CliMessage) Edit(configs ...CliConfigFunc) {
	for _, fn := range configs {
		fn(cfg)
	}
}

func WithSolve(solve string) CliConfigFunc {
	return func(cfg *CliMessage) {
		cfg.Solve = solve
	}
}

func WithDoc(doc string) CliConfigFunc {
	return func(cfg *CliMessage) {
		cfg.Doc = doc
	}
}

func (cfg *CliMessage) String() string {
	return ""
}

func (cfg CliMessage) ToStdout() {
	if _, err := os.Stdout.WriteString(cfg.String()); err != nil {
		log.Fatal("sla...")
	}
	setExitCode(cfg.Message.Level)
}

func (cfg *CliMessage) ToStderr() {
	if _, err := os.Stderr.WriteString(cfg.String()); err != nil {
		log.Fatal("sla...")
	}

	setExitCode(cfg.Message.Level)
}

func setExitCode(lvl Level) {
	if lvl == ERROR {
		os.Exit(1)
	}
}
