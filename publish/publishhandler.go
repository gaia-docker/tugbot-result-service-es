package publish

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
	body, err := getBodyReader(request)
	if err != nil {
		retStatus = http.StatusBadRequest
		log.Error(err)
	}
	params := request.URL.Query()
	publisher := JsonPublisher{}
	_, err = publisher.Publish(body, params.Get("docker.imagename"))
	if err != nil {
		retStatus = http.StatusInternalServerError
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
