package config

import "github.com/go-raptor/raptor/v2"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api",
			raptor.Scope("/v1",
				raptor.Route("POST", "/grammar/check", "GrammarController", "Check"),
			),
		),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
