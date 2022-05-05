package gonf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type gonf interface {
	Load(string) error
	Set(string) error
	Get(string) (string, error)
	GetKeys() []string
}

type config struct {
	conf map[string]string
}

func New(filePath string) (newConf config, err error) {
	newConf.conf = make(map[string]string)

	if len(filePath) > 0 {
		err = newConf.Load(filePath)
	}

	return
}

func (c config) Load(path string) (err error) {
	if strings.HasSuffix(path, ".json") {
		err = c.loadJson(path)
	} else if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		err = c.laodYaml(path)
	}
	return
}

func (c config) Set(key string, val string) (err error) {
	_, exists := c.conf[key]
	if exists {
		err = errors.New("Key already existed")
	}
	c.conf[key] = val
	return
}

func (c config) Get(key string) (ret string, err error) {
	ret, exists := c.conf[key]
	if exists == false {
		err = errors.New("No such key")
	}
	return
}

func (c config) GetAll() (ret map[string]string) {
	ret = make(map[string]string)
	for k, v := range c.conf {
		ret[k] = v
	}
	return ret
}

func (c config) GetKeys() (ret []string) {
	keys := make([]string, 0, len(c.conf))
	for k := range c.conf {
		keys = append(keys, k)
	}
	return keys
}

func (c config) loadJson(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &c.conf)
	return
}

func (c config) laodYaml(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, &c.conf)
	return
}
