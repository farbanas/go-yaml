package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func OpenFileRead(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	return data
}

func OpenFileWrite(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func ReadYaml(data []byte) interface{} {
	var out interface{}
	err := yaml.Unmarshal(data, &out)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	return out
}

func GetValue(yamlData interface{}, query string) {
	keys := strings.Split(query, ".")
	for _, key := range keys {
		yamlData = yamlData.(map[interface{}]interface{})[key]
	}
	yamlData = yamlData.([]interface{})[0].(string)
	fmt.Println(yamlData)
}
