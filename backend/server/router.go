package server

import (
	"github.com/gin-gonic/gin"

	swaggerdocs "my-playground/backend/server/swagger"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title my-playground (backend/server)
// @version 0.0.1
// @description "My Playground API"
//
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
//
func setup(handler *gin.Engine) *gin.Engine {
	swaggerdocs.SwaggerInfo.BasePath = "/my-playground-api"

	handler.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.PersistAuthorization(true),
		),
	)

	return handler
}
