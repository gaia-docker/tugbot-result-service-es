package elasticclient

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v3"
	"regexp"
)

const illegalCharactersRegexp = "[:,?<>/\\*?| ]"

type ESClient struct {
	client *elastic.Client
}

func NewESClient(url string) *ESClient {

	client, _ := elastic.NewClient(elastic.SetURL(url))

	return &ESClient{client}
}

func (esc *ESClient) CreateIndexIfNotExist(name string, mapping map[string]interface{}) (string, error) {

	validIndexName := esc.removeIllegalCharacters(name)
	exist, err := esc.client.IndexExists(validIndexName).Do()
	if err != nil {
		log.Errorf("Failed checking if index exist %s. %+v", validIndexName, err)
	}
	if !exist {
		_, err = esc.client.CreateIndex(validIndexName).Do()
		if err != nil {
			log.Errorf("Failed creating index %s. %+v", validIndexName, err)
		} else {
			log.Infof("Index <%s> created", validIndexName)
		}
		if mapping != nil {
			_, err = esc.client.PutMapping().Index(validIndexName).BodyJson(mapping).Do()
			if err != nil {
				log.Errorf("Failed creating mapping for index %s. %+v", validIndexName, err)
			} else {
				log.Infof("Created mapping for index <%s>", validIndexName)
			}
		}
	}

	return validIndexName, err
}

func (esc *ESClient) Index(indexName, docType string, doc interface{}) error {

	response, err := esc.client.Index().
		Index(indexName).
		Type(docType).
		BodyJson(doc).
		Refresh(true).
		Do()
	if err != nil {
		log.Errorf("Failed indexing document %+v for index %s. %+v", doc, indexName, err)
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
