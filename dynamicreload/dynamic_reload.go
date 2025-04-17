package dynamicreload

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Rizwan102003/dynamic-cms-go/content"
	"github.com/Rizwan102003/dynamic-cms-go/router"
	"github.com/gin-gonic/gin"
)

func RegisterNewContentType(c *gin.Context) {
	var req struct {
		Name   string            `json:"name"`
		Schema map[string]string `json:"schema"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := make([]reflect.StructField, 0)

	for fieldName, fieldType := range req.Schema {
		var typ reflect.Type
		switch fieldType {
		case "string":
			typ = reflect.TypeOf("")
		case "int":
			typ = reflect.TypeOf(0)
		case "float64":
			typ = reflect.TypeOf(0.0)
		case "bool":
			typ = reflect.TypeOf(true)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported type: " + fieldType})
			return
		}

		fields = append(fields, reflect.StructField{
			Name: capitalize(fieldName),
			Type: typ,
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, fieldName)),
		})
	}

	newStructType := reflect.StructOf(fields)

	// Register in global content registry
	content.RegisterContentType(req.Name, reflect.New(newStructType).Interface())

	// Generate new routes on the fly
	routerInstance := c.MustGet("router").(*gin.Engine)
	router.GenerateRoutesForType(routerInstance, req.Name, newStructType)

	c.JSON(http.StatusOK, gin.H{"message": "Content type '" + req.Name + "' registered successfully"})
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
