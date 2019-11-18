package main

const reamdeFile = `
## API GW Project

### Install Dependencies
You need the following dependencies to be able to build and deploy the api project:
* A working Pyhton installation
* awscli
* sam (serverless application model)
* create a new S3-bucket to store deployment files

**Install awscli and SAM**  
` + "```" + `bash
$ pip install --user awscli
$ pip install --user aws-sam-cli
` + "```" + `

### Build the Project
` + "```" + `bash
$ sam package --template-file apigw.yml  --output-template-file out.yaml --s3-bucket {Your-S3-bucket}
` + "```" + `

### Deploy the API
` + "```" + `bash
$ sam deploy --template-file ./out.yaml --stack-name your-api-project --capabilities CAPABILITY_IAM 
` + "```" + `


`
