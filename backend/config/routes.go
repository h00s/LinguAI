package config

import "github.com/go-raptor/raptor/v2"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api",
			raptor.Scope("/v1",
				raptor.Scope("/grammar",
					raptor.Route("POST", "/check", "GrammarController", "Check"),
				),
				raptor.Scope("/translator"),
				raptor.Scope("/languages",
					raptor.Route("GET", "/", "LanguagesController", "Index"),
				),
				raptor.Scope("/models",
					raptor.Route("GET", "/", "ModelsController", "All"),
				),
			),
		),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
