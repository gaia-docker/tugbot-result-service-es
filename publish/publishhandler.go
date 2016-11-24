package publish

import (
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
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
	body, err := getGZipBodyReader(request)
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

func (ph PublishHandler) HandleEvent(writer http.ResponseWriter, request *http.Request) {

	status := http.StatusOK
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		ph.esClient.Index("tugbot", "event", string(body))
	} else {
		status = http.StatusBadRequest
		log.Error("Failed to read request JSON body.", err)
	}
	writer.WriteHeader(status)
}

func getGZipBodyReader(request *http.Request) (io.ReadCloser, error) {

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
