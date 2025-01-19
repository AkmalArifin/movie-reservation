package middleware

import (
	"github.com/gin-contrib/cors"
)

func ConfigCORS() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	return config
}
