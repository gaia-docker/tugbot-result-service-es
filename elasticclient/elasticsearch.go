package elasticclient

import (
	log "github.com/Sirupsen/logrus"
	"github.com/olivere/elastic"
)

type ESClient struct {
	client *elastic.Client
}

func NewESClient() *ESClient {

	client, _ := elastic.NewClient()
	return &ESClient{client}
}

func (esc *ESClient) CreateIndexIfNotExist(name string) error {

	exist, err := esc.client.IndexExists(name).Do()
	if err != nil {
		log.Errorf("Failed checking if index exist %s. %+v", name, err)
	}
	if !exist {
		_, err = esc.client.CreateIndex(name).Do()
		if err != nil {
			log.Errorf("Failed creating index %s. %+v", name, err)
		} else {
			log.Debugf("Index %s created", name)
		}
	}

	return err
}

func (esc *ESClient) Index(indexName, doc string) error {

	response, err := esc.client.Index().
		Index(indexName).
		BodyString(doc).
		Do()
	if err != nil {
		log.Errorf("Failed indexing document %s for index %s. %+v", doc, indexName, err)
	} else {
		log.Debugf("Indexed document %s to index %s", response.Id, response.Index)
	}

	return err
}
