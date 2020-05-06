<p align="center"><a href="https://github.com/rogerwelin/alviss"><img src="logo.png" alt="alviss"></a></p>
<p align="center">
  <a href="https://goreportcard.com/badge/github.com/rogerwelin/alviss"><img src="https://goreportcard.com/badge/github.com/rogerwelin/alviss" alt="Go Report Card"></a>
  <a href="https://github.com/rogerwelin/alviss/blob/master/LICENSE"><img src="https://img.shields.io/github/license/rogerwelin/alviss" alt="License"></a>
  <a href="https://github.com/rogerwelin/alviss/releases"><img src="https://img.shields.io/github/v/release/rogerwelin/alviss.svg" alt="Current Release"></a>
  <a href="https://github.com/rogerwelin/alviss/blob/master/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/rogerwelin/alviss" alt="Go version"></a>
</p>


**Alviss** is a scaffolding project that let's you provision and deploy production ready serverless API:s in seconds on AWS using API Gateway and Lambda using your preferred programming language


Rationale
--------
Configuring API Gateway and Lambda using standard IaC tools like Terraform and Cloudformation is a very finicky, verbose and time consuming experience. Even with tools that are designed for serverless applications like *AWS SAM* and *serverless framework* can be hard and time consuming. Alviss is a scaffolder that takes care of generating the boilerplate using best practices and leaves you to tweak or modify the settings as you like. Use [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html) to finally deploy the project.

Support for *serverless framework* is upcoming


Installation
--------
Alviss is built in Go; meaning no runtime or dependencies to install, just grab a pre-built binary from the [GitHub Releases page](https://github.com/rogerwelin/alviss/releases). You can optionally put the **alviss** binary in your `PATH` so you can run alviss from any location.


Usage
--------

```bash
$ ./alviss new-api
NAME:
   alviss new-api - Generates a new api project

USAGE:
   alviss new-api [command options] [arguments...]

OPTIONS:
   -p value, --project-name value  name of your API project
   -t value, --api-type value      api type (only rest supported for now) (default: rest)
   -e value, --api-endpoint value  which endpoint type (either regional, edge or private) (default: regional)
   -l value, --language value      which language for lambda to be used (go, node, python, ruby) (default: node)
```

And example below shows how to generate a new serverless public api using node.js as target language for the Lambda function(s):

```bash
$ ./alviss new-api --project-name my-api-project --api-endpoint regional --language node
```

Then just follow the instructions on the screen. It's that simple!


Demo
--------

<img src="https://i.imgur.com/Zy8PG73.gif" />


Compliments
--------
Special thanks goes to [Axfood IT AB](https://www.axfood.se/) for letting me opensource this


