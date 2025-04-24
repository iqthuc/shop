package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"shop/internal/features/auth/core/dto"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignUpSuccess(t *testing.T) {
	t.Parallel()
	randomEmail := fmt.Sprintf("test-%d@example.com", time.Now().UnixNano())
	requestBody := dto.SignUpRequest{
		Email:    randomEmail,
		Password: "123456",
	}
	bodyJson, err := json.Marshal(requestBody)
	require.NoError(t, err)

	request := httptest.NewRequest(http.MethodPost, "/auth/sign-up", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := testApp.Server.Fiber.Test(request)
	defer func() {
		err := response.Body.Close()
		require.NoError(t, err)
	}()
	require.NoError(t, err)
	bytes, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	log.Println("Body:", string(bytes))

	responseBody := new(dto.SignUpResponse)
	err = json.Unmarshal(bytes, responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
