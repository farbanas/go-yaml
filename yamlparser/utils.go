package yamlparser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
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
	return out
}

func writeYaml(yamlData interface{}, f *os.File, path string) {
	out, err := yaml.Marshal(yamlData)
	if err != nil {
		log.Fatalf("Failed marshaling updated yaml data.")
	}
	_, err = f.Write(out)
	if err != nil {
		log.Fatalf("Failed writing updated yaml version to file %s.", path)
	}
}

func GetValue(yamlData interface{}, query string) interface{} {
	keys := strings.Split(query, ".")
	for _, key := range keys {
		var ok bool

		if intKey, err := strconv.Atoi(key); err == nil {
			var temp []interface{}
			temp, ok = yamlData.([]interface{})
			if ok {
				if len(temp) > intKey {
					yamlData = temp[intKey]
				} else {
					log.Fatalf("Error: index out of range [%d] for array of length %d", intKey, len(temp))
				}
			} else {
				log.Printf("Error: got array index, but current element is not an array.")
				log.Fatalf("Current element: %v", yamlData)
			}
		} else {
			temp, ok := yamlData.(map[interface{}]interface{})
			if ok {
				yamlData = temp[key]
			} else {
				log.Printf("Error: got key, but current element is not a map.")
				log.Fatalf("Current element: %v", yamlData)
			}
		}
	}
	return yamlData
}

func SetValue(yamlData interface{}, query string, val interface{}, path string) {
	data := yamlData.(map[interface{}]interface{})
	p := reflect.ValueOf(&data).Elem()

	keys := strings.Split(query, ".")

	for i, key := range keys {
		if p.Kind() == reflect.Map {
			mapKeys := p.MapKeys()
			for _, mapKey := range mapKeys {
				if mapKey.Interface().(string) == key {
					if i == len(keys)-1 {
						p.SetMapIndex(mapKey, reflect.ValueOf(val))
					}
					p = p.MapIndex(mapKey)
					v := reflect.ValueOf(p.Interface())
					p = reflect.New(v.Type()).Elem()
					p.Set(v)
				}
			}
		} else if p.Kind() == reflect.Slice {
			index, err := strconv.Atoi(key)
			if err != nil {
				log.Println("Current yaml level is array, but index not provided in query.")
				log.Fatalf("Current level: %v", p.Interface())
			}
			p = p.Index(index)
			if i == len(keys)-1 {
				p.Set(reflect.ValueOf(val))
			}
			v := reflect.ValueOf(p.Interface())
			p = reflect.New(v.Type()).Elem()
			p.Set(v)
		}
	}

	f := OpenFileWrite(path)
	writeYaml(yamlData, f, path)
	f.Close()
}
