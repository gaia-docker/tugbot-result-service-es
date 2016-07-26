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

func NewJsonPublisher() *JsonPublisher {

	return &JsonPublisher{elasticclient.NewESClient()}
}

func (jp *JsonPublisher) Publish(reader io.ReadCloser, indexNameSuffix string) (*string, error) {

	indexName := "tugbot_" + indexNameSuffix
	log.Infof("Going to publish results to index: %s", indexName)
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Failed reading stream: %+v", err)
		return nil, err
	}
	createdIndexName, err := jp.esClient.CreateIndexIfNotExist(indexName)
	if err != nil {
		return nil, err
	}

	return nil, jp.esClient.Index(createdIndexName, documentType, string(buffer))
}
