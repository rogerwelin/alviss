package main

import (
	"os"
	"text/template"
)

var (
	defaultAppPath          = "src/helloworld/app"
	allowedAPIProtocols     = []string{"rest", "websocket"}
	allowedRestAPIEndpoints = []string{"regional", "edge", "private"}
)

type TmplData struct {
	ApiProtocol        string
	ApiEndpoints       string
	LambdaFunctionName string
	ApiProjectName     string
}

type LanguageMapper struct {
	Files       []string
	TmplAppVar  string
	TmplDepsVar string
	AppPath     string
	DepsPath    string
}

// Usage to-do
var languages = map[string]LanguageMapper{
	"node": LanguageMapper{
		Files:       []string{"index.js", "package.json"},
		TmplAppVar:  "nodeFunction",
		TmplDepsVar: "packageJson",
		AppPath:     defaultAppPath,
		DepsPath:    defaultAppPath,
	},
	"java": LanguageMapper{
		Files:       []string{"App.java", "pom.xml"},
		TmplAppVar:  "",
		TmplDepsVar: "",
		AppPath:     defaultAppPath + "/com/api",
		DepsPath:    defaultAppPath,
	},
	"python": LanguageMapper{
		Files:       []string{"app.py", "requirements.txt"},
		TmplAppVar:  "",
		TmplDepsVar: "",
		AppPath:     defaultAppPath,
		DepsPath:    defaultAppPath,
	},
	"ruby": LanguageMapper{
		Files:       []string{"app.rb", "Gemfile"},
		TmplAppVar:  "",
		TmplDepsVar: "",
		AppPath:     defaultAppPath,
		DepsPath:    defaultAppPath,
	},
	"go": LanguageMapper{
		Files:       []string{"main.go", "go.mod"},
		TmplAppVar:  "",
		TmplDepsVar: "",
		AppPath:     defaultAppPath,
		DepsPath:    defaultAppPath,
	},
}

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

func (tmpl *TmplData) createFileFromTemplate(tmplVar string, path string, outName string) error {
	t := template.Must(template.New("").Parse(tmplVar))
	var file *os.File
	var err error

	if path != "" {
		file, err = os.Create(path + "/" + outName)
		if err != nil {
			return err
		}
	} else {
		file, err = os.Create(outName)
		if err != nil {
			return err
		}
	}
	err = t.Execute(file, tmpl)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	apiTmpl := &TmplData{
		ApiProtocol:        "rest",
		ApiEndpoints:       "regional",
		LambdaFunctionName: "helloworld",
		ApiProjectName:     "Hello-World-API",
	}

	err := apiTmpl.createFileFromTemplate(apiGWConf, "", "apigw.yml")
	if err != nil {
		panic(err)
	}

	err = apiTmpl.createFileFromTemplate(swagger, "", "swagger-api.yml")
	if err != nil {
		panic(err)
	}

	err = createDir("src/helloworld/app")
	if err != nil {
		panic(err)
	}

	err = apiTmpl.createFileFromTemplate(nodeFunction, "src/helloworld/app", "index.js")
	if err != nil {
		panic(err)
	}

	err = apiTmpl.createFileFromTemplate(packageJson, "src/helloworld/app", "package.json")
	if err != nil {
		panic(err)
	}

}
