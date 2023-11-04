package helpers

import "github.com/gin-contrib/cors"

var CorsConfig cors.Config

func CorsHeader() {
	// Configure CORS settings
	CorsConfig = cors.DefaultConfig()
	CorsConfig.AllowOrigins = []string{"*"}
	CorsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	CorsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "AuthenticationType", "SecurityType"}
	CorsConfig.ExposeHeaders = []string{"Content-Length"}
}
