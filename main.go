package main

import (
	"fmt"
	"go-boilerplate/utils"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"

	docs "go-boilerplate/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Boilerplate API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3050

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://dev-7iha82mtaab0qlu7.us.auth0.com/oauth/token
func main() {

	err := utils.LoadEnv()

	if err != nil {
		panic("env missing")
	}

	utils.SetupDatabase(utils.Envs.DSN)
	utils.AddValidators()

	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

	// TODO: only enable client to call
	// r.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowHeaders:     []string{"Authorization"},
	// }))
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	ApiRoutes(r)

	// Form errors
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
		opt.ValidatePrivateFields = true
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":" + utils.Envs.APP_PORT); err != nil {
		log.Fatalf("Unable to serve %v", err.Error())
	}

	fmt.Println("Serving at " + utils.Envs.APP_PORT)
}
