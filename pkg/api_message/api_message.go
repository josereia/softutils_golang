package messages

import (
	"fmt"
	"net/http"
)

/*
svc.Message.NotFound("sla...", methods)
svc.Message.BadRequest("sla...", methods)
*/
type ApiConfigFunc func(*ApiMessage)

type ApiMessage struct {
	HttpCode   int     `json:"http_code"`
	HttpStatus string  `json:"http_status"`
	Message    Message `json:"message,omitempty"`
}

func NewApiMessage(httpCode int, configs ...MessageFunc) ApiMessage {
	newMsg := Message{
		Description: "sla...",
	}

	for _, fn := range configs {
		fn(&newMsg)
	}

	return ApiMessage{
		HttpCode:   httpCode,
		HttpStatus: generateHttpStatusMsg(httpCode),
		Message:    newMsg,
	}
}

func (msg *ApiMessage) Update(configs ...ApiConfigFunc) {
	for _, fn := range configs {
		fn(msg)
	}
}

func (msg *ApiMessage) String() string {

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

func BadRequest(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusBadRequest
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func Conflict(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusConflict
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func Created(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusCreated
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func Forbidden(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusForbidden
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func InternalServerError(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusInternalServerError
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func NoContent(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusNoContent
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func NotFound(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusNotFound
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}

func Unauthorized(cfg *ApiMessage) {
	cfg.HttpCode = http.StatusUnauthorized
	cfg.HttpStatus = generateHttpStatusMsg(cfg.HttpCode)
}
