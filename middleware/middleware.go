package middleware

import (
	"../data"

	"github.com/Env-Co-Ltd/framinGo"
)

type Middleware struct {
	App    *framinGo.FraminGo
	Models data.Models
}
