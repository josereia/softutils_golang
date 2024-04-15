package messages

import (
	"fmt"
	"net/http"
)

type ApiConfigFunc func(*apiMessage)

type apiMessage struct {
	HttpCode   int     `json:"http_code"`
	HttpStatus string  `json:"http_status"`
	Timestamp  string  `json:"timestamp"`
	Message    Message `json:"message,omitempty"`
}

func NewApiMessage(httpCode int, msg string, messageFuncs ...MessageFunc) apiMessage {
	return setUpApiMessage(
		setUpMessage(msg, messageFuncs...), setHttpConfig(httpCode),
	)
}

func (msg *apiMessage) Update(configs ...ApiConfigFunc) {
	for _, fn := range configs {
		fn(msg)
	}
}

func (msg *apiMessage) String() string {

	if showMessageLvl == DEBUG && msg.Message.Stack != "" {
		return fmt.Sprintf(
			"[%s] >  %s\nstacktrace= %s",
			msg.HttpStatus,
			msg.Message.Description,
			msg.Message.Stack,
		)
	}

	return fmt.Sprintf(
		"[%s] >  %s",
		msg.HttpStatus,
		msg.Message.Description,
	)

}

func setHttpConfig(code int) ApiConfigFunc {
	return func(apiMsg *apiMessage) {
		apiMsg.HttpCode = code
		apiMsg.HttpStatus = generateHttpStatusMsg(code)
	}
}

func generateHttpStatusMsg(code int) string {
	var status string
	switch messageLanguage {
	case EN:
		status = http.StatusText(code)
	case PT:
		status = HttpStatusPt(code)
	}

	return fmt.Sprintf("%d %s", code, status)
}

func BadRequest(cfg *apiMessage) {
	setHttpConfig(http.StatusBadRequest)
}

func Conflict(cfg *apiMessage) {
	setHttpConfig(http.StatusConflict)
}

func Forbidden(cfg *apiMessage) {
	setHttpConfig(http.StatusForbidden)
}

func Internal(cfg *apiMessage) {
	setHttpConfig(http.StatusInternalServerError)
}

func NoContent(cfg *apiMessage) {
	setHttpConfig(http.StatusNoContent)
}

func NotFound(cfg *apiMessage) {
	setHttpConfig(http.StatusNotFound)
}

func Unauthorized(cfg *apiMessage) {
	setHttpConfig(http.StatusUnauthorized)
}
