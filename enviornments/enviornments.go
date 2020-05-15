package enviornments

import (
	"text/template"
	"io/ioutil"
	"bytes"
	"strings"
	"net/http"
	"log"
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
func BuildDockerfile(response GitHubResponse,dockerfile string, fileData FileData) string {
	index := search(response, dockerfile)

	if index != 1 {
		data, err := http.Get(response[index].DownloadUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer data.Body.Close()
		body, err := ioutil.ReadAll(data.Body)

		t := template.Must(template.New("dockerfile").Parse(string(body)))
	
		var templated bytes.Buffer
		t.Execute(&templated, fileData)
	
		return templated.String()
	}
	return ""
}

func search(arr GitHubResponse, str string) int {
	for index, item := range arr {
		if strings.ToLower(item.Name) == str {
			return index
		}
	}
	return -1
}