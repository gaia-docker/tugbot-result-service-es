package publish

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
	"io"
	"io/ioutil"
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
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Failed decoding stream into TestResult: %+v", err)
		return nil, err
	}
	createdIndexName, err := jp.esClient.CreateIndexIfNotExist(indexName)
	doc := string(buffer)
	log.Infof("Going to publish results to index: %s <%s>", createdIndexName, doc)
	if err != nil {
		return nil, err
	}

	return nil, jp.esClient.Index(createdIndexName, documentType, doc)
}
