package api

import (
	"go-rest-toolkit/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(Logger())
	r.Use(RateLimiter(10, 100))

	// Public routes
	r.POST("/login", LoginHandler)
	r.POST("/refresh", RefreshTokenHandler)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(auth.Authorize("admin", "user")) // now returns gin.HandlerFunc
	{
		authorized.GET("/profile", ProfileHandler)
	}

	return r
}
