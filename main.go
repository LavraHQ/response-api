package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lavrahq/response-api/config"
	"github.com/lavrahq/response-api/routes/auth"
	"github.com/lavrahq/response-api/routes/metadata"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		return
	}

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	requiresAuth := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.JwtTokenClaims{},
		SigningKey: []byte(config.Config.AuthJwtSecret.Key),
	})

	meta := e.Group("/api/metadata", requiresAuth)
	meta.GET("", metadata.GetMetadata)
	meta.GET("/special", metadata.GetSpecial)
	meta.GET("/:table", metadata.GetTableMetadata)

	e.Logger.Fatal(e.Start(":8100"))
}
