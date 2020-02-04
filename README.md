## Installation
To install `go-yaml` run the following command.
```
$ go get github.com/farbanas/go-yaml
```

Try running it in your terminal.
```
$ go-yaml
Error: subcommand has to be `get` or `set`.
Usage: go-yaml (get|set)
```

In case you get something like
```
$ go-yaml
go-yaml: command not found
```
please check that your `$GOPATH/bin` is included in your PATH. You can set it for the current shell by running:
```
$ export PATH:${GOPATH:-~/go}/bin
```
To save that configuration, put it in your `~/.bashrc` or `~/.zshrc`.

TODO: examples, usage
