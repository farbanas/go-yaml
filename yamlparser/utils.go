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
	var dataMap map[interface{}]interface{} = nil
	var dataArray []interface{} = nil

	keys := strings.Split(query, ".")

	for _, key := range keys {
		dataMap, dataArray = mapOrArray(yamlData)
		if dataMap != nil {
			if s, ok := dataMap[key]; ok {
				yamlData = s
			} else {
				log.Fatalf("Key %s doesn't exist.", key)
			}
		} else if dataArray != nil {
			index, err := strconv.Atoi(key)
			if err != nil {
				log.Println("Current yaml level is array, but index not provided in query.")
				log.Fatalf("Current level: %v", dataArray)
			}
			yamlData = dataArray[index]
		}
	}
	return yamlData
}

func GetValueReflect(yamlData interface{}, query string) interface{} {
	p := reflect.ValueOf(&yamlData).Elem()

	keys := strings.Split(query, ".")

	for _, key := range keys {
		if p.Kind() == reflect.Map {
			mapKey := reflect.ValueOf(key)
			p = p.MapIndex(reflect.ValueOf(mapKey))
			v := reflect.ValueOf(p.Interface())
			p = reflect.New(v.Type()).Elem()
			p.Set(v)
		} else if p.Kind() == reflect.Slice {
			index, err := strconv.Atoi(key)
			if err != nil {
				log.Println("Current yaml level is array, but index not provided in query.")
				log.Fatalf("Current level: %v", p.Interface())
			}
			p = p.Index(index)
			v := reflect.ValueOf(p.Interface())
			p = reflect.New(v.Type()).Elem()
			p.Set(v)
		}
	}
	return p
}

func SetValue(yamlData interface{}, query string, val interface{}, path string) {
	var dataMap map[interface{}]interface{} = nil
	var dataArray []interface{} = nil
	tempIntf := yamlData

	keys := strings.Split(query, ".")

	for i, key := range keys {
		dataMap, dataArray = mapOrArray(tempIntf)
		if dataMap != nil {
			if s, ok := dataMap[key]; ok {
				if i == len(keys)-1 {
					dataMap[key] = val
				} else {
					tempIntf = s
				}
			} else {
				log.Fatalf("Key %s doesn't exist.", key)
			}
		} else if dataArray != nil {
			index, err := strconv.Atoi(key)
			if err != nil {
				log.Println("Current yaml level is array, but index not provided in query.")
				log.Fatalf("Current level: %v", dataArray)
			}
			if i == len(keys)-1 {
				dataArray[index] = val
			} else {
				tempIntf = dataArray[index]
			}
		}
	}

	f := OpenFileWrite(path)
	writeYaml(yamlData, f, path)
	f.Close()
}

func SetValueReflect(yamlData interface{}, query string, val interface{}, path string) {
	p := reflect.ValueOf(&yamlData).Elem()

	keys := strings.Split(query, ".")

	for i, key := range keys {
		if p.Kind() == reflect.Map {
			mapKey := reflect.ValueOf(key)
			if i == len(keys)-1 {
				p.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(val))
			} else {
				p = p.MapIndex(reflect.ValueOf(mapKey))
				v := reflect.ValueOf(p.Interface())
				p = reflect.New(v.Type()).Elem()
				p.Set(v)
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
			} else {
				v := reflect.ValueOf(p.Interface())
				p = reflect.New(v.Type()).Elem()
				p.Set(v)
			}
		}
	}

	f := OpenFileWrite(path)
	writeYaml(yamlData, f, path)
	f.Close()
}

func mapOrArray(data interface{}) (dataMap map[interface{}]interface{}, dataArray []interface{}) {
	var ok bool
	dataMap, ok = data.(map[interface{}]interface{})
	if !ok {
		dataArray, ok = data.([]interface{})
		if !ok {
			return nil, nil
		}
	}
	return
}
