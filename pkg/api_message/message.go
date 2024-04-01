package messages

import "errors"

// import "github.com/pkg/errors"

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
	Description string `json:"description,omitempty"`
	Stack       string `json:"stack,omitempty"`
	Level       Level  `json:"level,omitempty"`
	RootError   error  `json:"root_error,omitempty"`
}

func WithCustomDescription(description string) MessageFunc {
	return func(msg *Message) {
		msg.Description = description
	}
}

func WithRootError(err error) MessageFunc {
	return func(msg *Message) {
		msg.RootError = err
		msg.Description = err.Error()
		msg.Stack = errors.Unwrap(err).Error()
	}
}

func (msg *Message) Edit(configs ...MessageFunc) {
	for _, fn := range configs {
		fn(msg)
	}
}
