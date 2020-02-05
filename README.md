# YAGS
<!-- vim-markdown-toc -->
Yags is a simple YAML getter/setter written in GO. 

Current features are:
- get any value from any valid yaml file
- set a value for a key in yaml file

## Table of contents
* [Installation](#installation)
* [Usage](#usage)
	* [Get](#get)
	* [Set](#set)
* [Upcoming](#upcoming)


## Installation
To install `yags` run the following command.
```
$ go get github.com/farbanas/yags
```

Try running it in your terminal.
```
$ yags
Error: subcommand has to be `get` or `set`.
Usage: yags (get|set)
```

In case you get something like
```
$ yags
yags: command not found
```
please check that your `$GOPATH/bin` is included in your PATH. You can set it for the current shell by running:
```
$ export PATH:${GOPATH:-~/go}/bin
```
To save that configuration, put it in your `~/.bashrc` or `~/.zshrc`.

## Usage
Yags has two subcommands that it supports, `get` and `set`.
```
$ yags
Error: subcommand has to be `get` or `set`.
Usage: yags (get|set)
```
### Get
The `get` subcommand is used for getting a value from yaml file. 
```
$ yags get 
Error: filePath cannot be empty!
Error: query cannot be empty!

Usage of yags get:
  -filePath string
        Path to yaml file.
  -query string 
        Query for the value to get. Query should be in the dot format, for example if you want to set the value of a yaml map entry that is on the third level,
        your query would look something like 'first.second.third'. It also supports array indexes (indexes are 0-indexed). In the case that you have an array,
        your query would look something like 'first.second.2.third'.
```
As you can see from the command usage, it requires two parameters, path to the yaml file from which you want to get the value
and a query that specifies which value you want. The first parameter is straightforward, but I will explain query a bit more.

Query is a dot separated collection of keys that are parents to the value you want to get. For example, you have a yaml like this
(this is a randomly generated yaml btw):
```yaml
potatoes:
- percent: true
  soil:
    engineer: programmer
    through:
    - rabbit
    - grass
    - station
    tip: mind
  stiff: tight
- false
- true
satellites: record
worry: deer
```
If you wanted to get the value of engineer, your query would look like `potatoes.0.soil.engineer`. The only thing that should 
be clarified in this string is the `0`. It represents the index of the yaml array. We are taking the first element (0-indexed).
At that point, you could imagine that you now have a yaml like this:
```yaml
percent: true
soil:
  engineer: programmer
  through:
  - rabbit
  - grass
  - station
  tip: mind
stiff: tight
```
which is why we need `soil` next and `engineer` after that. If you saved that yaml into a file named `example.yaml` and ran
`yags` on it, you would get something like this:
```shell script
$ yags get -filePath example.yaml -query potatoes.0.soil.engineer
programmer
``` 
Different example, where we want to get `grass`:
```shell script
$ yags get -filePath example.yaml -query potatoes.0.soil.through.1
grass
```
### Set
The set subcommand is similar to the get subcommand, it just needs one parameter, the value that you want to set.
```
$ yags set
Error: filePath cannot be empty!
Error: key cannot be empty!
Error: value cannot be empty!

Usage of yags get:
  -filePath string
        Path to yaml file.
  -query string 
        Query for the value to set. Query should be in the dot format, for example if you want to set the value of a yaml map entry that is on the third level,
        your query would look something like 'first.second.third'. It also supports array indexes (indexes are 0-indexed). In the case that you have an array,
        your query would look something like 'first.second.2.third'.
  -value string
        Value that you want to set.
```
Querying works the same as for the get subcommand. 

Let's take the same yaml file as for the get subcommand.
```yaml
potatoes:
- percent: true
  soil:
    engineer: programmer
    through:
    - rabbit
    - grass
    - station
    tip: mind
  stiff: tight
- false
- true
satellites: record
worry: deer
```
Let's say we want to change the value `mind` which is associated to key `tip` to `please`. It would look like this:
```shell script
$ yags set -filePath example.yaml -query potatoes.0.soil.tip -value please
$ echo "$?"
0
```
In this case yags doesn't output anything, but it finished successfully which can be seen if you inspect the exit code.
You will now have a yaml file looking like this:
```yaml
potatoes:
- percent: true
  soil:
    engineer: programmer
    through:
    - rabbit
    - grass
    - station
    tip: please
  stiff: tight
- false
- true
satellites: record
worry: deer
```

## Upcoming
- [ ] When setting, create a key if it does not exist (should be switchable with command line option)
- [ ] Set complex values instead of only strings
- [ ] Better logging
- [ ] Rewrite set to not use reflect


