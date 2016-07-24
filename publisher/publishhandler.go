package publisher

import (
	"compress/gzip"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func Handle(writer http.ResponseWriter, request *http.Request) {

	retStatus := http.StatusOK
	_, err := getBodyReader(request)
	if err != nil {
		retStatus = http.StatusBadRequest
		log.Error(err)
	}
	writer.WriteHeader(retStatus)
}

func getBodyReader(request *http.Request) (io.ReadCloser, error) {

	body := request.Body
	if body == nil {
		return nil, errors.New("Empty request body")
	}
	contentType := request.Header.Get("Content-Type")
	if strings.Contains(contentType, "gzip") {
		return gzip.NewReader(body), nil
	}

	return body, nil
}
