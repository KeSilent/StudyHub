package initialize

import (
	_ "github.com/KeSilent/study-hub/server/source/example"
	_ "github.com/KeSilent/study-hub/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
