package publish

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
	"io"
	"io/ioutil"
)

const typeTestCase = "testcase"

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
		log.Errorf("Failed reading stream %+v", err)
		return nil, err
	}
	createdIndexName, err := jp.esClient.CreateIndexIfNotExist(indexName)
	if err != nil {
		return nil, err
	}
	doc := string(buffer)
	log.Infof("Going to publish results to index: %s <%s>", createdIndexName, doc)

	return nil, jp.esClient.Index(createdIndexName, typeTestCase, doc)
}
