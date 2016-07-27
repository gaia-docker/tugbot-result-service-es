package publish

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
	"github.com/gaia-docker/tugbot-result-service-es/model"
	"io"
)

const documentType = "testcase"

type JsonPublisher struct {
	esClient *elasticclient.ESClient
}

func NewJsonPublisher(esClient *elasticclient.ESClient) *JsonPublisher {

	return &JsonPublisher{esClient}
}

func (jp *JsonPublisher) Publish(reader io.ReadCloser, indexNameSuffix string) (*string, error) {

	indexName := "tugbot_" + indexNameSuffix
	var testResult model.TestResult
	err := json.NewDecoder(reader).Decode(&testResult)
	if err != nil {
		log.Errorf("Failed decoding stream into TestResult: %+v", err)
		return nil, err
	}
	createdIndexName, err := jp.esClient.CreateIndexIfNotExist(indexName, model.GetTestResultMapping())
	log.Infof("Going to publish results to index: %s <%+v>", createdIndexName, testResult)
	if err != nil {
		return nil, err
	}

	return nil, jp.esClient.Index(createdIndexName, documentType, testResult)
}
