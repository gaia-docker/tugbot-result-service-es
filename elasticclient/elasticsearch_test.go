package elasticclient

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveIllegalCharacters_replace(t *testing.T) {

	publisher := NewESClient()
	indexName := publisher.removeIllegalCharacters("gaia,integartion<tests>:latest")
	log.Infof("index name: %s", indexName)
	assert.Equal(t, "gaia_integartion_tests__latest", indexName)
}

func TestRemoveIllegalCharacters(t *testing.T) {

	publisher := NewESClient()
	indexName := publisher.removeIllegalCharacters("gaia-integartion-tests.latest")
	log.Infof("index name: %s", indexName)
	assert.Equal(t, indexName, indexName)
}
