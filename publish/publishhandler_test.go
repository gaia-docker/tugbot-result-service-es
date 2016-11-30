package publish

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleEvent(t *testing.T) {
	called := false
	handler := NewPublishHandler(httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			called = true
			body, err := ioutil.ReadAll(request.Body)
			assert.NoError(t, err)
			assert.True(t, strings.Contains(string(body),
				fmt.Sprintf(`"%s":`, TugbotTimeFieldName)))
			writer.Write([]byte(fmt.Sprintf(`{"_index":"%s","_created":"true"}`,
				strings.Split(request.URL.RequestURI(), "/")[1])))
		})).URL)
	request, err := http.NewRequest("POST", "/events", strings.NewReader(`
		{"Status":"create",
		"ID":"04f10323b0d9ef32411024333d7fd15e65ea856ffab95425bffc65dbd63dc106",
		"From":"perl","Type":"container","Action":"create"}`))
	assert.NoError(t, err)
	response := httptest.NewRecorder()
	handler.HandleEvent(response, request)
	assert.True(t, called)
	assert.Equal(t, http.StatusOK, response.Code)
}
