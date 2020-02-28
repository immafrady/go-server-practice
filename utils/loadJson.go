package utils

import (
	"encoding/json"
	"errors"
	_struct "fradyspace.com/go-server-practice/struct"
	"log"
	"os"
	"path"
)

var base string
var err error

type pathMap map[_struct.Path]string

var pathConfig pathMap

func init() {
	base, err = os.Getwd()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	pathConfig = loadPathConfig()
}

func LoadConfig(fileName _struct.Env) (ret _struct.Config) {
	file, err := os.Open(formatPath(string(fileName)))
	if err != nil {
		log.Fatalf(err.Error())
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&ret)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return
}

func loadPathConfig() (ret pathMap) {
	ret = make(pathMap)
	file, err := os.Open(formatPath("path"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&ret)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return
}

func GetPathConfig(path _struct.Path) (string, error) {
	if v, ok := pathConfig[path]; ok {
		return base + v, nil
	} else {
		return "", errors.New("unable to locate")
	}
}

func formatPath(fileName string) string {
	return path.Join(base, "config", fileName+".json")
}
