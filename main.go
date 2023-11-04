package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manishbadgotra/authentication/controllers"
	"github.com/manishbadgotra/authentication/helpers"
	"github.com/manishbadgotra/authentication/middlewares"
	util "github.com/manishbadgotra/authentication/utility"
)

func init() {

	// Init config
	helpers.InitConfig()
	helpers.CorsHeader()
	util.ConnectDB()
}

func main() {
	router := gin.Default

	open := router().Group("/api/user")
	secure := router().Group("/api/user").Use(middlewares.ValidateToken)

	router().Use(cors.New(helpers.CorsConfig))
	open.Use(cors.New(helpers.CorsConfig))
	secure.Use(cors.New(helpers.CorsConfig))

	open.POST("/signup", controllers.Signup)
	secure.GET("/checkstatus", controllers.CheckStatus)

	router().Run(helpers.GetString("JwtSecretKey"))
}
