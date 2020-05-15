package enviornments

import (
	"text/template"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
	"strings"
)

type Env struct {
	name string
	location string
	id string
	container string
}
type envFile struct {
	Enviornments []enviornment `json:"enviornments"`
}
type enviornment struct { 
	Name string `json:"name"`
	DockerName string `json:"dockerName"`
	DockerFile string `json:"dockerFile"`
}
type FileData struct {
	Port int
	Main string
	MainExt string
}
func CreateEnv(name string) Env {
	return Env{}
}
func BuildDockerfile(dockerfile string, fileData FileData) {
	path := "/home/andrew/dev/personal/projects/1ProductAWeek/week2/code/enviornments/enviornments.json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	
	var file envFile
	err = json.Unmarshal(data, &file)
	if err != nil {
		fmt.Print(err)
	}
	index := search(file.Enviornments, dockerfile)

	if index != -1 {
		t := template.Must(template.New("dockerfile").Parse(file.Enviornments[index].DockerFile))
	
		var templated bytes.Buffer
		t.Execute(&templated, fileData)
	
		fmt.Println(templated.String())
	}
}
// func main() {

// }

func search(arr []enviornment, str string) int {
	for index, item := range arr {
		if strings.ToLower(item.Name) == str {
			return index
		}
	}
	return -1
}