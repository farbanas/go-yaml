package main

import (
	"flag"
	"fmt"
	"os"
)

func flagParser() (string, string, string) {
	action := flag.Arg(0)
	fp := flag.String("filePath", "", "Path to yaml file.")
	key := flag.String("key", "", "Yaml key for the value to get/set. Ex. first.second will get "+
		"the value of key second that is a subkey of first.")
	var val string
	if action == "set" {
		val = *flag.String("value", "", "Value that you want to set.")
	}
	flag.Parse()
	if *fp == "" || *key == "" {
		if *fp == "" {
			fmt.Println("Path to yaml file cannot be empty!")
		}
		if *key == "" {
			fmt.Println("Yaml key cannot be empty!")
		}
		flag.PrintDefaults()
		os.Exit(1)
	}
	return *fp, *key, val
}

func main() {
	fp, key, _ := flagParser()
	data := OpenFileRead(fp)
	yamlData := ReadYaml(data)
	GetValue(yamlData, key)
}
