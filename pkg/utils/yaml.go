package uitls

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// DocsConfig config data structure for config yaml file parser
type DocsConfig struct {
	docsYamlConfig DocsYamlConfig
	fileName       string
}

func NewDocsConfig(fileName string) *DocsConfig {
	return &DocsConfig{
		fileName: fileName,
	}
}

// DocsYamlConfig corresponding golang structure for versions yaml file
type DocsYamlConfig struct {
	Versions []struct {
		Ver   string `yaml:"ver"`
		Repos []struct {
			Name    string `yaml:"name"`
			URL     string `yaml:"url"`
			TagName string `yaml:"tagName"`
		} `yaml:"repos"`
	} `yaml:"versions"`
}

// GetDocsYamlConfig
func (c *DocsConfig) GetDocsYamlConfig() DocsYamlConfig {
	return c.docsYamlConfig
}

// Parse parse yaml config file
func (c *DocsConfig) Parse() error {
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &c.docsYamlConfig)
	return err
}
