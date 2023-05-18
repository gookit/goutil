// Package structs Provide some extends util functions for struct. eg: tag parse, struct init, value set
package structs

// IsExported field name on struct
func IsExported(fieldName string) bool {
	return fieldName[0] >= 'A' && fieldName[0] <= 'Z'
}

// IsUnexported field name on struct
func IsUnexported(fieldName string) bool {
	return fieldName[0] >= 'a' && fieldName[0] <= 'z'
}
