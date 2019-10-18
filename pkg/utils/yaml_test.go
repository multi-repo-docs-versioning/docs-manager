package uitls

import (
	"testing"

	"gotest.tools/assert"
)

func TestYamlParser(t *testing.T) {

	config := NewDocsConfig("../../configs/versions.yml")

	err := config.Parse()
	assert.Equal(t, err, nil)
}
