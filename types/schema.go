package types

type ContentType struct {
	Name   string
	Fields map[string]string // fieldName -> fieldType
}

func NewContentType(fields map[string]string) *ContentType {
	return &ContentType{
		Fields: fields,
	}
}
