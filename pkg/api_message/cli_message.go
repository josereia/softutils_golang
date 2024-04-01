package messages

import (
	"encoding/json"
	"fmt"
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
	Title   string  `json:"title"`
	Doc     string  `json:"doc,omitempty"`
	Solve   string  `json:"solve,omitempty"`
	Message Message `json:"message"`
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

func (msg *CliMessage) Edit(configs ...CliConfigFunc) {
	for _, fn := range configs {
		fn(msg)
	}
}

func WithSolve(solve string) CliConfigFunc {
	return func(msg *CliMessage) {
		msg.Solve = solve
	}
}

func WithDoc(doc string) CliConfigFunc {
	return func(msg *CliMessage) {
		msg.Doc = doc
	}
}

func (msg *CliMessage) String() string {
	formattedMsg := fmt.Sprintf(
		"[%s] %s\n%s\n", msg.Message.RootError, msg.Title, msg.Message.Description,
	)

	if msg.Solve != "" {
		formattedMsg += fmt.Sprintf("%s\n", msg.Solve)
	}

	if msg.Doc != "" {
		formattedMsg += fmt.Sprint(msg.Doc)
	}

	return formattedMsg

}

func (msg *CliConfigFunc) BindJson() string {
	jsonData, err := json.MarshalIndent(msg, " ", "  ")
	if err != nil {
		panic("Error on marshal json on softutils in module CliMessage")
	}

	return string(jsonData)
}

func (msg CliMessage) ToStdout() {
	if _, err := os.Stdout.WriteString(msg.String()); err != nil {
		log.Fatal("sla...")
	}
	setExitCode(msg.Message.Level)
}

func (msg *CliMessage) ToStderr() {
	if _, err := os.Stderr.WriteString(msg.String()); err != nil {
		log.Fatal("sla...")
	}

	setExitCode(msg.Message.Level)
}

func setExitCode(lvl Level) {
	if lvl == ERROR {
		os.Exit(1)
	}
}
