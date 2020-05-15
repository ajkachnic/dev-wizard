package cli

import (
	"flag"
	"fmt"

	"github.com/ajkachnic/dev-wizard/enviornments"
	"strings"
	"os"
	"path/filepath"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

var env string
var port int = 0
var main string
var dir string
var list bool


func init() {
	flag.StringVar(&env, "env", "", `The enviornment you would like to configure`)
	flag.StringVar(&env, "e", "", `The enviornment you would like to configure`)

	flag.IntVar(&port, "port", 0, `The port to expose for docker`)
	flag.IntVar(&port, "p", 0, `The port to expose for docker`)

	flag.StringVar(&main, "main", "", `The main file of the project`)
	flag.StringVar(&main, "m", "", `The main file of the project`)

	flag.StringVar(&dir, "dir", ".", `The directory to create the file in`)
	flag.StringVar(&dir, "d", ".", `The directory to create the file in`)

	flag.BoolVar(&list, "list", false, "Lists all the available enviornments")
	flag.BoolVar(&list, "l", false, "Lists all the available enviornments")

	flag.Parse()
}

func Run() {
	url := "https://api.github.com/repos/ajkachnic/dev-wizard/contents/templates"
	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result enviornments.GitHubResponse
	json.Unmarshal(body, &result)

	var envs []string
	for _, v := range result {
		envs = append(envs, v.Name)
	}

	if list {
		fmt.Println("Envs:")
		for _, v := range envs {
			fmt.Println("- " + v)
		}
	} else {
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
			file := enviornments.BuildDockerfile(result, env, fileData)
			if file != "" {
				path, err := filepath.Abs("./")
				if err != nil {
					log.Fatal(err)
				}
				perm := os.FileMode(0644)
				if dir == "." || dir == "" {
					path, err := filepath.Abs("./")
					err = ioutil.WriteFile(path + "/Dockerfile", []byte(file), perm)
					if err != nil {
						log.Fatal(err)
					}
				} else {
					if dir[0] == '/'  || dir[0] == '~' {
						err = ioutil.WriteFile(dir + "/Dockerfile", []byte(file), perm)
						if err != nil {
							log.Fatal(err)
						}
					} else {
						err = ioutil.WriteFile(path + "/" + dir + "/Dockerfile", []byte(file), perm)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			} else {
				fmt.Println("Invalid enviornment")
			}
		} else {
			fmt.Println("Invalid enviornment")
		}
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