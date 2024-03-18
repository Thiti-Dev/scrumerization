package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testUser struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name" jsonschema:"required"`
	Username string `json:"username" jsonschema:"required,minLength=4"`
}

// go test -timeout 30s -run ^TestValidateJsonSchemaFromStruct$ github.com/Thiti-Dev/scrumerization-core-service/pkg/validation --count=1 -v
func TestValidateJsonSchemaFromStruct(t *testing.T) {
	validationErr, valid := ValidateJsonSchemaFromStruct(testUser{Username: "1"})
	assert.False(t, valid)
	assert.NotEmpty(t, validationErr)

	assert.Equal(t, 1, len(validationErr.FieldErrors["missing_properties"])) // name is missing
}
