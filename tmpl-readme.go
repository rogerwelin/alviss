package main

const reamdeFile = `
## API GW Project

### Set up the Environment
To create the API GW project you will need an AWS account and generate API keys for your IAM user. Run *aws configure* after installing the awscli. Then you you will need to create a new S3 bucket that will be used for storing deployment files.

### Install Dependencies
You need the following dependencies to be able to build and deploy the api project:
* A working Pyhton installation
* awscli
* sam (serverless application model)

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

Go to the AWS console > Cloudformation. Make sure the Cloudformations stack finishes. Take a look at the output to get the URL of your newly created API project. Either curl the address at the /hello endpoint or run the endpoint directly in the API Gateway console.


`
