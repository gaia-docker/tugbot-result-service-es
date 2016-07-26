package publish

import (
	"compress/gzip"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
	"io"
	"net/http"
	"strings"
)

type PublishHandler struct {
	esClient *elasticclient.ESClient
}

// NewUploadHandler creates UploadHandler instance
func NewPublishHandler(esUrl string) *PublishHandler {

	return &PublishHandler{elasticclient.NewESClient(esUrl)}
}

func (ph PublishHandler) Handle(writer http.ResponseWriter, request *http.Request) {

	retStatus := http.StatusOK
	body, err := getBodyReader(request)
	if err != nil {
		retStatus = http.StatusBadRequest
		log.Error(err)
	}
	params := request.URL.Query()
	publisher := NewJsonPublisher(ph.esClient)
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
		return gzip.NewReader(body)
	}

	return body, nil
}
