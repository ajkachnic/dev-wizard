package enviornments

import (
	"text/template"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
	"strings"
	"net/http"
)

type Env struct {
	name string
	location string
	id string
	container string
}

type FileData struct {
	Port int
	Main string
	MainExt string
}
type GitHubResponse []GitHubResponseFile
type GitHubResponseFile struct {
	Name string `json:"name"`
	DownloadUrl string `json:"download_url"`
}

func CreateEnv(name string) Env {
	return Env{}
}
func BuildDockerfile(dockerfile string, fileData FileData) {
	url := "https://api.github.com/repos/ajkachnic/dev-wizard/contents/templates"
	data, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)

	var result GitHubResponse
	json.Unmarshal(body, &result)

	index := search(result, dockerfile)

	if index != 1 {
		data, err = http.Get(result[index].DownloadUrl)
		if err != nil {
			fmt.Print(err)
		}
		defer data.Body.Close()
		body, err = ioutil.ReadAll(data.Body)

		t := template.Must(template.New("dockerfile").Parse(string(body)))
	
		var templated bytes.Buffer
		t.Execute(&templated, fileData)
	
		fmt.Println(templated.String())
	}
}

func search(arr GitHubResponse, str string) int {
	for index, item := range arr {
		if strings.ToLower(item.Name) == str {
			return index
		}
	}
	return -1
}