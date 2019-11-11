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
}
