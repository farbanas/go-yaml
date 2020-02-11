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
	keyGet := getFlagSet.String("query", "", "Query for the value to get. Query should be in the dot format, "+
		"for example if you want to set the value of a yaml map entry that is on the third level,\n"+
		"your query would look something like 'first.second.third'. It also supports array indexes (indexes are 0-indexed)."+
		" In the case that you have an array,\nyour query would look something like 'first.second.2.third'.")

	fpSet := setFlagSet.String("filePath", "", "Path to yaml file.")
	keySet := setFlagSet.String("query", "", "Query for the value to set. Query should be in the dot format, "+
		"for example if you want to set the value of a yaml map entry that is on the third level,\n"+
		"your query would look something like 'first.second.third'. It also supports array indexes (indexes are 0-indexed)."+
		" In the case that you have an array,\nyour query would look something like 'first.second.2.third'.")
	valSet := setFlagSet.String("value", "", "Value that you want to set, it will be put as string.")

	parseSubcommand(getFlagSet, setFlagSet)
	if os.Args[1] == "get" {
		parseRequiredFlags(getFlagSet, []string{"filePath", "query"}, fpGet, keyGet)
		return *fpGet, *keyGet, ""
	} else {
		parseRequiredFlags(setFlagSet, []string{"filePath", "query", "value"}, fpSet, keySet, valSet)
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
		if len(flagNames) == 3 {
			fmt.Println("Usage of yags set:")
		} else {
			fmt.Println("Usage of yags get:")
		}
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
		fmt.Println(result)
	} else if os.Args[1] == "set" {
		yamlparser.SetValue(yamlData, query, val, fp)
	}
}
