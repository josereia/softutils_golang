package messages

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiConfigFunc func(*ApiMessage)

type ApiMessage struct {
	HttpCode   int     `json:"http_code"`
	HttpStatus string  `json:"http_status"`
	Timestamp  string  `json:"timestamp"`
	Message    Message `json:"message,omitempty"`
}

func NewApiMessage(httpCode int, msg string, messageFuncs ...MessageFunc) ApiMessage {
	return ApiMessage{}
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

// Set up and send apiError to ctx.JSON.
func (msg *ApiMessage) ToResponseJSON(ctx *gin.Context) {
	fmt.Println(">>> HTTP ERROR:\n ", msg)
	ctx.JSON(msg.HttpCode, msg)
}

func setHttpConfig(code int, apiMsg *ApiMessage) {
	apiMsg.HttpCode = code
	apiMsg.HttpStatus = generateHttpStatusMsg(code)
}

func generateHttpStatusMsg(code int) string {
	var status string
	switch messageLanguage {
	case EN:
		status = http.StatusText(code)
	case PT:
		status = HttpStatusPt(code)
	}

	return fmt.Sprintf("[%d] %s", code, status)
}

func BadRequest(cfg *ApiMessage) {
	setHttpConfig(http.StatusBadRequest, cfg)
}

func Conflict(cfg *ApiMessage) {
	setHttpConfig(http.StatusConflict, cfg)
}

func Forbidden(cfg *ApiMessage) {
	setHttpConfig(http.StatusForbidden, cfg)
}

func Internal(cfg *ApiMessage) {
	setHttpConfig(http.StatusInternalServerError, cfg)
}

func NoContent(cfg *ApiMessage) {
	setHttpConfig(http.StatusNoContent, cfg)
}

func NotFound(cfg *ApiMessage) {
	setHttpConfig(http.StatusNotFound, cfg)
}

func Unauthorized(cfg *ApiMessage) {
	setHttpConfig(http.StatusUnauthorized, cfg)
}
