package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Validate(t *testing.T, test TestCase, rr *httptest.ResponseRecorder) {
	if rr.Code != 0 {
		assert.Equal(t, test.Expected.StatusCode, rr.Code)
	}

	if test.Expected.BodyPart != "" {
		assert.Contains(
			t,
			rr.Body.String(),
			test.Expected.BodyPart,
			fmt.Sprintf("expecting '%v' in %v", test.Expected.BodyPart, rr.Body.String()),
		)
	}
	if len(test.Expected.BodyParts) > 0 {
		for _, expectedText := range test.Expected.BodyParts {
			assert.Contains(t, rr.Body.String(), expectedText)
		}
	}
	if test.Expected.BodyPartMissing != "" {
		assert.NotContains(t, rr.Body.String(), test.Expected.BodyPartMissing)
	}

	if len(test.Expected.BodyPartsMissing) > 0 {
		for _, expectedText := range test.Expected.BodyPartsMissing {
			assert.NotContains(t, rr.Body.String(), expectedText)
		}
	}

	if test.Expected.Callback != nil {
		test.Expected.Callback(t)
	}
}
func RunTest(t *testing.T, test TestCase, ts *TestServer) {
	reqJson, err := json.Marshal(test.RequestBody)
	if err != nil {
		t.Fatalf("Error creating request body: %v", err)
	}

	req, err := http.NewRequest(test.Request.Method, test.Request.Url, bytes.NewBuffer(reqJson))
	if err != nil {
		t.Fatalf("unable to create the request %v", err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rr := httptest.NewRecorder()

	ts.S.Echo.ServeHTTP(rr, req)
	// Validate results
	Validate(t, test, rr)
}
