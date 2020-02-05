# YAGS
Yags is a simple YAML getter/setter written in GO. 

Current features are:
- get any value from any valid yaml file
- set a value for a key in yaml file

TODO:
- [ ] When setting, create a key if it does not exist (should be switchable with command line option)
- [ ] Set complex values instead of only strings
- [ ] Better logging
- [ ] Rewrite set to not use reflect

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
```
$ yags

```
