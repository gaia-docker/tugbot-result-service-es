package main

import (
	"compress/gzip"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func publish(writer http.ResponseWriter, request *http.Request) {

	retStatus := http.StatusOK
	_, err := getGzip(request)
	if err != nil {
		retStatus = http.StatusBadRequest
		log.Error(err)
	}
	writer.WriteHeader(retStatus)
}

func getGzip(request *http.Request) (io.ReadCloser, error) {

	body := request.Body
	if body == nil {
		return nil, errors.New("Empty request body")
	}
	if !strings.Contains(request.Header.Get("Content-Type"), "gzip") {
		return nil, errors.New("Content-Type should be gzip")
	}

	return gzip.NewReader(body)
}
