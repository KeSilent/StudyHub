package response

import (
	"github.com/KeSilent/study-hub/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
