package publish

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
)

const TugbotTimeFieldName = "tugbot-time"

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
		ph.esClient.Index("tugbot", "event", addTimeToJson(body))
	} else {
		status = http.StatusBadRequest
		log.Error("Failed to read request JSON body.", err)
	}
	writer.WriteHeader(status)
}

func addTimeToJson(json []byte) string {

	event := string(json)
	event = event[1:len(event)]

	return fmt.Sprintf(`{"%s":"%s", %s`, TugbotTimeFieldName, time.Now().Format(time.RFC3339), event)
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
