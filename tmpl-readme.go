package main

const reamdeFile = `
## API GW Project

### Set up the Environment
To create the API GW project you will need an AWS account and generate API keys for your IAM user. Run *aws configure* after installing the awscli.

### Install Dependencies
You need the following dependencies to be able to build and deploy the api project:
* A working Python installation
* awscli
* sam (serverless application model)
* docker (optional, only if you plan to run the api locally)

**Install awscli and SAM**  
` + "```" + `bash
$ pip install --user awscli
$ pip install --user aws-sam-cli
` + "```" + `

### Build all the Lambda functions
` + "```" + `bash
$ sam build --template-file apigw.yml
` + "```" + `

### Package & Deploy the Project
` + "```" + `bash
$ sam deploy --guided
` + "```" + `

The --guided flag will set up all needed IAM roles and a S3 bucket and finally will set upp all needed AWS resources and deploy the code.

Check the Outputs from the command line and grab the URL. Now do a curl against URL/helloworld
{{ if and (eq .APIEndpoints "private") }}
However since this is a private api, you need to do the curl from within your VPC
{{ end }}
`
