package request

import (
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
