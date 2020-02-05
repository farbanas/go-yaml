package main

import (
	"flag"
	"fmt"
	"github.com/farbanas/yags/yamlparser"
	"log"
	"os"
)

func flagParser() (string, string, string) {
	getFlagSet := flag.NewFlagSet("get", flag.ExitOnError)
	setFlagSet := flag.NewFlagSet("set", flag.ExitOnError)

	fpGet := getFlagSet.String("filePath", "", "Path to yaml file.")
	keyGet := getFlagSet.String("key", "", "Yaml key for the value to get. Ex. first.second will get "+
		"the value of key second that is a subkey of first.")

	fpSet := setFlagSet.String("filePath", "", "Path to yaml file.")
	keySet := setFlagSet.String("key", "", "Yaml key for the value to set. Ex. first.second will set "+
		"the value of key second that is a subkey of first.")
	valSet := setFlagSet.String("value", "", "Value that you want to set.")

	parseSubcommand(getFlagSet, setFlagSet)
	if os.Args[1] == "get" {
		parseRequiredFlags(getFlagSet, []string{"filePath", "key"}, fpGet, keyGet)
		return *fpGet, *keyGet, ""
	} else {
		parseRequiredFlags(setFlagSet, []string{"filePath", "key", "value"}, fpSet, keySet, valSet)
		return *fpSet, *keySet, *valSet
	}
}

func parseRequiredFlags(flagSet *flag.FlagSet, flagNames []string, flags ...*string) {
	exit := false
	for i, f := range flags {
		if *f == "" {
			fmt.Printf("Error: %s cannot be empty!\n", flagNames[i])
			exit = true
		}
	}
	if exit {
		fmt.Println()
		fmt.Println("Usage of yags:")
		flagSet.PrintDefaults()
		os.Exit(2)
	}
}

func parseSubcommand(getFlagSet *flag.FlagSet, setFlagSet *flag.FlagSet) {
	if len(os.Args) < 2 {
		fmt.Println("Error: subcommand has to be `get` or `set`.")
		fmt.Println("Usage: yags (get|set)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		err := getFlagSet.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "set":
		err := setFlagSet.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "default":
		fmt.Println("Error: subcommand has to be `get` or `set`.")
		fmt.Println("Usage: yags (get|set)")
		os.Exit(1)
	}
}

func main() {
	fp, query, val := flagParser()
	data := yamlparser.OpenFileRead(fp)
	yamlData := yamlparser.ReadYaml(data)
	if os.Args[1] == "get" {
		result := yamlparser.GetValue(yamlData, query)
		fmt.Printf("Result: %v\n", result)
	} else if os.Args[1] == "set" {
		yamlparser.SetValue(yamlData, query, val, fp)
	}
}
