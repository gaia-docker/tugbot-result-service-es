package publish

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gaia-docker/tugbot-result-service-es/elasticclient"
	"io"
	"io/ioutil"
)

type JsonPublisher struct {
	esClient *elasticclient.ESClient
}

func (jp JsonPublisher) Publish(reader io.ReadCloser, indexNameSuffix string) (*string, error) {

	indexName := "tugbot_" + indexNameSuffix
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Failed reading stream: %+v", err)
		return nil, err
	}
	err = jp.esClient.CreateIndexIfNotExist(indexName)
	if err != nil {
		return nil, err
	}

	return nil, jp.esClient.Index(indexName, string(buffer))
}
