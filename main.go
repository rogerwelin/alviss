package main

import (
	"os"
	"text/template"
)

type TmplData struct {
	ApiProtocol        string
	ApiEndpoints       string
	LambdaFunctionName string
	ApiProjectName     string
}

var allowedAPIProtocols = []string{"rest", "websocket"}
var allowedRestAPIEndpoints = []string{"regional", "edge", "private"}

func createDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func contains(slice []string, contains string) bool {
	for _, item := range slice {
		if contains == item {
			return true
		}
	}
	return false
}

func main() {

	apiTmpl := TmplData{
		ApiProtocol:        "rest",
		ApiEndpoints:       "regional",
		LambdaFunctionName: "helloworld",
		ApiProjectName:     "Hello-World-API",
	}

	t := template.Must(template.New("apigw").Parse(apiGWConf))
	file, err := os.Create("apigw.yml")
	if err != err {
		panic(err)
	}
	err = t.Execute(file, apiTmpl)
	if err != nil {
		panic(err)
	}

	err = createDir("src/helloworld/app")
	if err != nil {
		panic(err)
	}
}
