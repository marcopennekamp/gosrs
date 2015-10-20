package core
import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Data struct {
		Path string
	}
}

func ReadConfig(path string) (*Config, error) {
	configString, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = yaml.Unmarshal(configString, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (conf *Config) GetFilePath(dataPath string) string {
	return filepath.Join(conf.Data.Path, dataPath)
}