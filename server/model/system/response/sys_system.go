package response

import "github.com/KeSilent/study-hub/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
