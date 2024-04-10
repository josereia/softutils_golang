package messages

import "time"

func NewBadRequest(msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), BadRequest)
}

func NewNotFound(msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), NotFound)
}

func NewConflict(msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Conflict)
}

func NewUnauthorized(msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Unauthorized)
}

func NewNoContent(msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), NoContent)
}

func NewForbidden(msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(setUpMessage(msg, messageFuncs...), Forbidden)
}

func setUpApiMessage(message Message, apiFuncs ...ApiConfigFunc) apiMessage {
	var apiMsg apiMessage

	InternalServerError(&apiMsg)

	for _, fn := range apiFuncs {
		fn(&apiMsg)
	}

	apiMsg.Timestamp = time.Now().Format(time.DateTime)
	apiMsg.Message = message

	return apiMsg
}

func setUpMessage(msg string, messageFuncs ...MessageFunc) Message {
	var message Message

	message.Description = msg

	for _, fn := range messageFuncs {
		fn(&message)
	}

	return message
}
