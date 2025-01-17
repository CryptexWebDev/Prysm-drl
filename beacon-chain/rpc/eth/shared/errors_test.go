package shared

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/lookup"
	"github.com/Dorol-Chain/Prysm-drl/v5/network/httputil"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/assert"
	"github.com/pkg/errors"
)

// TestWriteStateFetchError tests the WriteStateFetchError function
// to ensure that the correct error message and code are written to the response
// as an expected JSON format.
func TestWriteStateFetchError(t *testing.T) {
	cases := []struct {
		err             error
		expectedMessage string
		expectedCode    int
	}{
		{
			err:             &lookup.StateNotFoundError{},
			expectedMessage: "State not found",
			expectedCode:    http.StatusNotFound,
		},
		{
			err:             &lookup.StateIdParseError{},
			expectedMessage: "Invalid state ID",
			expectedCode:    http.StatusBadRequest,
		},
		{
			err:             errors.New("state not found"),
			expectedMessage: "Could not get state",
			expectedCode:    http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		writer := httptest.NewRecorder()
		WriteStateFetchError(writer, c.err)

		assert.Equal(t, c.expectedCode, writer.Code, "incorrect status code")
		assert.StringContains(t, c.expectedMessage, writer.Body.String(), "incorrect error message")

		e := &httputil.DefaultJsonError{}
		assert.NoError(t, json.Unmarshal(writer.Body.Bytes(), e), "failed to unmarshal response")
	}
}
