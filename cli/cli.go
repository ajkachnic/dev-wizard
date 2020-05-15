package cli

import (
	"flag"
	"fmt"

	"github.com/ajkachnic/dev-wizard/enviornments"
	"strings"
)

var env string
var port int = 0
var main string

var envs = []string{
	"node",
	"python",
	"go",
	"ruby",
	"php",
	"java",
}

func init() {
	flag.StringVar(&env, "env", "", `The enviornment you would like to configure
	Available options are
	- node
	- python
	- go
	- ruby
	- php
	- java
	`)
	flag.StringVar(&env, "e", "", `The enviornment you would like to configure
	Available options are
	- node
	- python
	- go
	- ruby
	- php
	- java
	`)

	flag.IntVar(&port, "port", 0, `The port to expose for docker`)
	flag.IntVar(&port, "p", 0, `The port to expose for docker`)

	flag.StringVar(&main, "main", "", `The main file of the project`)


	flag.Parse()
}

func Run() {
	if env == "" {
		fmt.Println("No enviornment specified. See --help")
	} else if main == "" {
		fmt.Println("No main specified. See --help")
	} else {
		if contains(envs, env) {
			fileData := enviornments.FileData{
				Port: port,
				Main: main,
			}
			if env == "java" {
				fileData.MainExt = main
				main = strings.TrimSuffix(main, ".java")
				fileData.Main = main
			}
			enviornments.BuildDockerfile(env, fileData)
		} else {
			fmt.Println("Invalid enviornment")
		}
	}
}
func contains(arr []string, str string) bool {
	for _, a := range arr {
		 if a == str {
				return true
		 }
	}
	return false
}
