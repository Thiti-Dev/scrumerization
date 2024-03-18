package tokenizer

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// go test -timeout 30s -run ^TestCreateToken$ github.com/Thiti-Dev/scrumerization-core-service/pkg/tokenizer -v --count=1
func TestCreateToken(t *testing.T) {
	payload, err := NewPayload(uuid.New(), "test", time.Hour)
	assert.NoError(t, err)
	token := CreateToken(payload, []byte("asdasd"))
	assert.NotEmpty(t, token)

	// Verifying the created token
	decryptedPayload, err := VerifyToken(token, []byte("asdasd"))
	assert.NoError(t, err)
	assert.NotEmpty(t, decryptedPayload)
}
