package main

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/fatih/color"
)

var (
	defaultAppPath = "/src/helloworld/app"
)

type tmplData struct {
	APIProtocol        string
	APIEndpoints       string
	LambdaFunctionName string
	APIProjectName     string
	Language           string
}

type languageMapper struct {
	AppFile     string
	DepsFile    string
	TmplAppVar  string
	TmplDepsVar string
	AppPath     string
	DepsPath    string
}

func initMap(projectName string) map[string]languageMapper {

	appPath := projectName + defaultAppPath

	var languages = map[string]languageMapper{
		"node": languageMapper{
			AppFile:     "index.js",
			DepsFile:    "package.json",
			TmplAppVar:  nodeFunction,
			TmplDepsVar: packageJson,
			AppPath:     appPath,
			DepsPath:    appPath,
		},
		"python": languageMapper{
			AppFile:     "app.py",
			DepsFile:    "requirements.txt",
			TmplAppVar:  pythonFunction,
			TmplDepsVar: requirementsFile,
			AppPath:     appPath,
			DepsPath:    appPath,
		},
		"ruby": languageMapper{
			AppFile:     "app.rb",
			DepsFile:    "Gemfile",
			TmplAppVar:  rubyFunction,
			TmplDepsVar: gemFile,
			AppPath:     appPath,
			DepsPath:    appPath,
		},
		"go": languageMapper{
			AppFile:     "main.go",
			DepsFile:    "go.mod",
			TmplAppVar:  goFunction,
			TmplDepsVar: goMod,
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

func createFileFromStruct(languageSpec languageMapper) error {
	// not pretty, fix later
	apa := &tmplData{}

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

func (tmpl *tmplData) createFileFromTemplate(tmplVar string, path string, outName string) error {
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

func (tmpl *tmplData) bootstrapAPI() error {
	col := color.New(color.FgCyan).Add(color.Underline)
	mg := color.New(color.FgGreen)
	fmt.Printf("\n\U0001f3c1  ")
	col.Printf("Bootstrapping API GW project: %s with Lambda\n\n", tmpl.APIProjectName)
	time.Sleep(350 * time.Millisecond)

	// create top dir
	createDir(tmpl.APIProjectName)

	// create readme
	mg.Printf("\u2705  Writing out README\n")
	time.Sleep(350 * time.Millisecond)
	err := tmpl.createFileFromTemplate(reamdeFile, tmpl.APIProjectName, "README.md")
	if err != nil {
		panic(err)
	}

	// create apigw sam/cf
	mg.Printf("\u2705  Writing out CF/SAM config\n")
	time.Sleep(350 * time.Millisecond)
	err = tmpl.createFileFromTemplate(apiGWConf, tmpl.APIProjectName, "apigw.yml")
	if err != nil {
		panic(err)
	}

	// create swagger file
	mg.Printf("\u2705  Writing out Swagger config\n")
	time.Sleep(350 * time.Millisecond)
	err = tmpl.createFileFromTemplate(swagger, tmpl.APIProjectName, "swagger-api.yml")
	if err != nil {
		return err
	}

	languageMap := initMap(tmpl.APIProjectName)

	mg.Printf("\u2705  Writing out %s lambda to: %s\n\n", tmpl.Language, tmpl.APIProjectName+defaultAppPath)
	err = createFileFromStruct(languageMap[tmpl.Language])
	if err != nil {
		return err
	}

	green := color.New(color.FgGreen).SprintFunc()
	consoleStdOut := "Success! Created API GW project at %s\nInside that directory," +
		"you can run several commands:\n\n\t%s\n\t\t" +
		"installs dependencies for all the Lambda source code\n\t%s\n\t\t" +
		"creates a zip of your code and dependencies " +
		"and uploads it to S3\n\t%s\n\t\t" +
		"deploys the specified CloudFormation/SAM template by creating and then " +
		"executing a change set\n\n" +
		"However I recommend taking a look at the README file first\n\n"

	//fmt.Printf("Success! Created API GW project at %s\nInside that directory, you can run several commands:\n\n\t%s\n\t\tcreates a zip of your code and dependencies and uploads it to S3\n\t%s\n\t\tdeploys the specified CloudFormation/SAM template by creating and then executing a change set\n\nHowever I recommend taking a look at the README file first\n\n", green(tmpl.APIProjectName+"/"), green("sam package --template-file apigw.yml  --output-template-file out.yaml --s3-bucket Your-S3-bucket"), green("aws cloudformation deploy --template-file ./out.yaml --stack-name my-api-stack --capabilities CAPABILITY_IAM"))
	fmt.Printf(consoleStdOut, green(tmpl.APIProjectName+"/"),
		green("sam build --template-file apigw.yml"),
		green("sam package --template-file apigw.yml  --output-template-file out.yaml --s3-bucket Your-S3-bucket"),
		green("aws cloudformation deploy --template-file ./out.yaml --stack-name my-api-stack --capabilities CAPABILITY_IAM"))

	return nil
}
