package middleware

import (
	"${APP_NAME}/data"

	"github.com/Env-Co-Ltd/framinGo"
)

type Middleware struct {
	App    *framinGo.FraminGo
	Models data.Models
}
