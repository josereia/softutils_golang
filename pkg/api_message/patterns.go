package messages

import "time"

func NewBadRequest(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), BadRequest)
}

func NewNotFound(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), NotFound)
}

func NewConflict(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Conflict)
}

func NewUnauthorized(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Unauthorized)
}

func NewNoContent(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), NoContent)
}

func NewForbidden(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Forbidden)
}

func NewInternal(msg string, messageFuncs ...MessageFunc) ApiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Internal)
}

func setUpApiMessage(message Message, apiFuncs ...ApiConfigFunc) ApiMessage {
	var apiMsg ApiMessage

	Internal(&apiMsg)

	for _, fn := range apiFuncs {
		fn(&apiMsg)
	}

	apiMsg.Timestamp = time.Now().Format(time.DateTime)
	apiMsg.Message = message

	return apiMsg
}

func setUpMessage(msg string, messageFuncs ...MessageFunc) Message {
	var message Message

	message.Title = msg

	for _, fn := range messageFuncs {
		fn(&message)
	}

	return message
}
