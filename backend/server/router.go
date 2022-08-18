package server

import (
	"github.com/gin-gonic/gin"

	swaggerdocs "my-playground/backend/server/swagger"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Playground (backend/server/router.go)
// @version 1.0.0
// @description "My Playground API Service"
//
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
//
func SetupHandler(handler *gin.Engine) *gin.Engine {
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
