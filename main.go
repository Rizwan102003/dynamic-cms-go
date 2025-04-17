package main

import (
	"github.com/Rizwan102003/dynamic-cms-go/content"
	"github.com/Rizwan102003/dynamic-cms-go/dynamicreload"
	"github.com/Rizwan102003/dynamic-cms-go/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Provide access to router in context for dynamic reload
	r.Use(func(c *gin.Context) {
		c.Set("router", r)
		c.Next()
	})

	// Route to dynamically register new content types
	r.POST("/register-content", func(c *gin.Context) {
		c.Set("router", r)
		dynamicreload.RegisterNewContentType(c)
	})

	// Generate routes for all existing registered content types
	for name, ct := range content.GetContentTypes() {
		router.GenerateRoutesForType(r, name, ct.Type) // Use ct.Type instead of ct
	}

	r.Run(":8080")
}
