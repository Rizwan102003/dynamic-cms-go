package content

import "reflect"

// ContentType holds both the type (schema) and a sample instance
type ContentType struct {
	Type    reflect.Type
	Example interface{}
}

// ContentRegistry maps content type name to its definition
var ContentRegistry = make(map[string]ContentType)

// RegisterContentType adds a new content type to the registry
func RegisterContentType(name string, model interface{}) {
	ContentRegistry[name] = ContentType{
		Type:    reflect.TypeOf(model),
		Example: model,
	}
}

// GetContentTypes returns the full content registry
func GetContentTypes() map[string]ContentType {
	return ContentRegistry
}

// GetContentType returns a single content type by name
func GetContentType(name string) (ContentType, bool) {
	ct, exists := ContentRegistry[name]
	return ct, exists
}
