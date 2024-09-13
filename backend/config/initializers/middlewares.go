package initializers

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/linguai/app/middlewares"
)

func Middlewares() raptor.Middlewares {
	return raptor.Middlewares{
		&middlewares.AuthMiddleware{},
	}
}
