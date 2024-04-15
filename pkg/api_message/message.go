package messages

import (
	"strings"

	"github.com/pkg/errors"
)

var showMessageLvl Level
var messageLanguage Language

func SetShowMessageLvl(lvl Level) {
	showMessageLvl = lvl
}

func SetMessageLanguage(lang Language) {
	messageLanguage = lang
}

type MessageFunc func(*Message)

type Message struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Stack       string `json:"stack,omitempty"`
	Level       Level  `json:"level,omitempty"`
	Error       string `json:"root_error,omitempty"`
}

func WithCustomDescription(description string) MessageFunc {
	return func(msg *Message) {
		msg.Description = description
	}
}

func WithError(err error) MessageFunc {
	return func(msg *Message) {
		msg.Description = strings.ToTitle(errors.Cause(err).Error())

		rootError := errors.Unwrap(err)

		if rootError != nil {
			msg.Stack = errors.Unwrap(err).Error()
		}
	}
}

func (msg *Message) Edit(configs ...MessageFunc) {
	for _, fn := range configs {
		fn(msg)
	}
}
