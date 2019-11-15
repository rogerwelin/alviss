package main

import (
	"os"
	"text/template"
)

var (
	defaultAppPath = "/src/helloworld/app"
)

type TmplData struct {
	ApiProtocol        string
	ApiEndpoints       string
	LambdaFunctionName string
	ApiProjectName     string
	Language           string
}

type LanguageMapper struct {
	AppFile     string
	DepsFile    string
	TmplAppVar  string
	TmplDepsVar string
	AppPath     string
	DepsPath    string
}

func initMap(projectName string) map[string]LanguageMapper {

	appPath := projectName + defaultAppPath

	var languages = map[string]LanguageMapper{
		"node": LanguageMapper{
			AppFile:     "index.js",
			DepsFile:    "package.json",
			TmplAppVar:  nodeFunction,
			TmplDepsVar: packageJson,
			AppPath:     appPath,
			DepsPath:    appPath,
		},
		"java": LanguageMapper{
			AppFile:     "App.java",
			DepsFile:    "pom.xml",
			TmplAppVar:  "",
			TmplDepsVar: "",
			AppPath:     appPath + "/com/api",
			DepsPath:    appPath,
		},
		"python": LanguageMapper{
			AppFile:     "app.py",
			DepsFile:    "requirements.txt",
			TmplAppVar:  "",
			TmplDepsVar: "",
			AppPath:     appPath,
			DepsPath:    appPath,
		},
		"ruby": LanguageMapper{
			AppFile:     "app.rb",
			DepsFile:    "Gemfile",
			TmplAppVar:  "",
			TmplDepsVar: "",
			AppPath:     appPath,
			DepsPath:    appPath,
		},
		"go": LanguageMapper{
			AppFile:     "main.go",
			DepsFile:    "go.mod",
			TmplAppVar:  "",
			TmplDepsVar: "",
			AppPath:     appPath,
			DepsPath:    appPath,
		},
	}
	return languages
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

func createFileFromStruct(languageSpec LanguageMapper) error {
	// not pretty, fix later
	apa := &TmplData{}

	err := createDir(languageSpec.AppPath)
	if err != nil {
		return err
	}
	err = apa.createFileFromTemplate(languageSpec.TmplAppVar, languageSpec.AppPath, languageSpec.AppFile)
	if err != nil {
		return err
	}
	err = apa.createFileFromTemplate(languageSpec.TmplDepsVar, languageSpec.DepsPath, languageSpec.DepsFile)
	if err != nil {
		return err
	}
	return nil
}

func (tmpl *TmplData) createFileFromTemplate(tmplVar string, path string, outName string) error {
	t := template.Must(template.New("").Parse(tmplVar))
	var file *os.File
	var err error

	file, err = os.Create(path + "/" + outName)
	if err != nil {
		return err
	}

	err = t.Execute(file, tmpl)
	if err != nil {
		return err
	}
	return nil
}

func (tmpl *TmplData) bootstrapAPI() error {

	// create top dir
	createDir(tmpl.ApiProjectName)

	// create apigw sam/cf
	err := tmpl.createFileFromTemplate(apiGWConf, tmpl.ApiProjectName, "apigw.yml")
	if err != nil {
		panic(err)
	}

	// create swagger file
	err = tmpl.createFileFromTemplate(swagger, tmpl.ApiProjectName, "swagger-api.yml")
	if err != nil {
		return err
	}

	languageMap := initMap(tmpl.ApiProjectName)

	err = createFileFromStruct(languageMap[tmpl.Language])
	if err != nil {
		return err
	}
	return nil
}

/*func main() {

	apiTmpl := &TmplData{
		ApiProtocol:        "rest",
		ApiEndpoints:       "regional",
		LambdaFunctionName: "helloworld",
		ApiProjectName:     "Hello-World-API",
		Language:           "node",
	}

	err := apiTmpl.createFileFromTemplate(apiGWConf, "", "apigw.yml")
	if err != nil {
		panic(err)
	}

	err = apiTmpl.createFileFromTemplate(swagger, "", "swagger-api.yml")
	if err != nil {
		panic(err)
	}

	err = createFileFromStruct(languages[apiTmpl.Language])
	if err != nil {
		panic(err)
	}

}
*/
