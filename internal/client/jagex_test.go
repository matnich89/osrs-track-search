package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJagexClient_GetHighScores(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name         string
		responseCode int
		expectError  bool
	}{
		{
			name:         "SuccessfulResponse",
			responseCode: http.StatusOK,
			expectError:  false,
		},
		{
			name:         "ErrorResponse",
			responseCode: http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.responseCode)
			}))
			defer server.Close()

			client, err := NewJagexClient(server.URL)
			assert.NoError(t, err, "NewJagexClient() should not return an error")

			_, err = client.GetHighScores("testCharacter", Normal)
			if tc.expectError {
				assert.Error(t, err, "GetHighScores() should return an error for this test case")
			} else {
				assert.NoError(t, err, "GetHighScores() should not return an error for this test case")
			}
		})
	}
}
