package elasticclient

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v3"
	"regexp"
	"strings"
)

const illegalCharactersRegexp = "[:,?<>/\\*?| ]"

type ESClient struct {
	client *elastic.Client
}

func NewESClient(url string) *ESClient {

	client, _ := elastic.NewSimpleClient(elastic.SetURL(url))

	return &ESClient{client}
}

func (esc *ESClient) CreateIndexIfNotExist(name string) (string, error) {

	validIndexName := esc.removeIllegalCharacters(name)
	exist, err := esc.client.IndexExists(validIndexName).Do()
	if err != nil {
		log.Errorf("Failed checking if index exist %s. %+v", validIndexName, err)
	}
	if !exist {
		_, err = esc.client.CreateIndex(validIndexName).Do()
		// ignore index_already_exists_exception error, probably race condition
		if err != nil && !strings.Contains(err.Error(), "index_already_exists_exception") {
			log.Errorf("Failed creating index %s. %+v", validIndexName, err)
		} else {
			err = nil
			log.Infof("Index <%s> created", validIndexName)
		}
	}

	return validIndexName, err
}

func (esc *ESClient) Index(indexName, docType, doc string) error {

	response, err := esc.client.Index().
		Index(indexName).
		Type(docType).
		BodyString(doc).
		Refresh(true).
		Do()
	if err != nil {
		log.Errorf("Failed indexing document %s for index %s. %+v", doc, indexName, err)
	} else {
		log.Infof("Indexed document %s to index %s", response.Id, response.Index)
	}

	return err
}

func (esc *ESClient) removeIllegalCharacters(indexName string) string {

	match, _ := regexp.MatchString(illegalCharactersRegexp, indexName)
	if !match {
		return indexName
	}
	r, _ := regexp.Compile(illegalCharactersRegexp)

	return r.ReplaceAllString(indexName, "_")
}
