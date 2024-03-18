package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	jsonSchemaReflect "github.com/invopop/jsonschema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type ValidationError struct {
	FieldErrors map[string][]string
}

// Specifically built to satisfy the gqlerror's extension type
func (e *ValidationError) FieldErrorsIntoArbitaryMapInterface() map[string]interface{} {
	convertedMap := make(map[string]interface{})
	for key, value := range e.FieldErrors {
		convertedMap[key] = interface{}(value)
	}
	return convertedMap
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// Extract property's names from string -> `missing properties: 'name', 'username', 'and', 'so', 'on'`
func extractMissingProperties(s string) []string {
	// Define a regular expression pattern to match property names
	pattern := `'([^']+)'`
	re := regexp.MustCompile(pattern)

	// Find all matches in the string
	matches := re.FindAllString(s, -1)

	// Extract property names from matches
	var properties []string
	for _, match := range matches {
		// Remove surrounding single quotes and add to properties slice
		properties = append(properties, match[1:len(match)-1])
	}

	return properties
}

func ValidateJsonSchemaFromStruct(baseStruct interface{}) (ValidationError, bool) {
	generatedJsonSchema := jsonSchemaReflect.Reflect(baseStruct)
	marshalledGeneratedJsonSchema, _ := json.Marshal(generatedJsonSchema)

	compiler := jsonschema.NewCompiler()
	err := compiler.AddResource("schema.json", bytes.NewReader(marshalledGeneratedJsonSchema))
	if err != nil {
		log.Fatalf("Failed to compile schema: %s", err)
	}
	schemaObj, err := compiler.Compile("schema.json")
	if err != nil {
		log.Fatalf("Failed to compile schema: %s", err)
	}

	inputAsJson, _ := json.Marshal(baseStruct)

	var inputAsMap map[string]interface{}
	json.Unmarshal(inputAsJson, &inputAsMap)

	// Remove the key with empty value
	// To sustain the missing-property error
	for k, v := range inputAsMap {
		reflectValue := reflect.ValueOf(v)
		if isEmptyValue(reflectValue) {
			delete(inputAsMap, k)
		}
	}

	err = schemaObj.Validate(inputAsMap)
	if err != nil {
		formattedError := make(map[string][]string)
		switch err := err.(type) {
		case jsonschema.InvalidJSONTypeError:
			// This would never happen unless we unmarshal the wrong format
		case *jsonschema.ValidationError:
			for _, validationErr := range err.BasicOutput().Errors {
				// Format the error message to include the path and the error message
				// fmt.Printf("ERROR: %s\n", validationErr.Error)
				// fmt.Printf("ERROR Detail: %s\n", validationErr.InstanceLocation)

				if strings.Contains(validationErr.Error, "doesn't validate with") {
					formattedErrorItems := formattedError["failed_causes"]
					errors := append(formattedErrorItems, validationErr.Error)
					formattedError["failed_causes"] = errors
				} else if strings.Contains(validationErr.Error, "missing properties") {
					formattedError["missing_properties"] = extractMissingProperties(validationErr.Error)
				} else if len(validationErr.InstanceLocation) > 0 {
					formattedErrorItems := formattedError["validation_errors"]
					errors := append(formattedErrorItems, fmt.Sprintf("%s -> %s", validationErr.InstanceLocation, validationErr.Error))
					formattedError["validation_errors"] = errors
				}
			}
			return ValidationError{FieldErrors: formattedError}, false
		default:
			log.Fatalf("got %v. want InvalidJSONTypeErr", err)
		}
	}

	return ValidationError{}, true
}
