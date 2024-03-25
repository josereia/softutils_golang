package messages

var emptyMsg = &ApiMessage{}

func (msg *ApiMessage) clearMsg() {
	*msg = *emptyMsg
}

func (msg *ApiMessage) NotFound(message string, configs ...MessageFunc) {
	msg.clearMsg()
	NotFound(msg)
	msg.Message.Description = message
	msg.Message.Edit(configs...)
}

func (msg *ApiMessage) BadRequest(message string, configs ...MessageFunc) {
	msg.clearMsg()
	BadRequest(msg)
	msg.Message.Description = message
	msg.Message.Edit(configs...)
}
