package router

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GenerateRoutesForType(r *gin.Engine, name string, typ reflect.Type) {
	basePath := "/" + name

	// GET all
	r.GET(basePath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GET all " + name})
	})

	// GET by ID
	r.GET(basePath+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "GET " + name + " with ID: " + id})
	})

	// POST new
	r.POST(basePath, func(c *gin.Context) {
		instance := reflect.New(typ).Interface()
		if err := c.ShouldBindJSON(&instance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Created new " + name, "data": instance})
	})

	// PUT by ID
	r.PUT(basePath+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		instance := reflect.New(typ).Interface()
		if err := c.ShouldBindJSON(&instance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Updated " + name + " with ID: " + id, "data": instance})
	})

	// DELETE by ID
	r.DELETE(basePath+"/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": "Deleted " + name + " with ID: " + id})
	})
}
