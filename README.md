<p align="center"><a href="https://github.com/rogerwelin/alviss"><img src="logo.png" alt="alviss"></a></p>
<p align="center">
  <a href="https://goreportcard.com/badge/github.com/rogerwelin/alviss"><img src="https://goreportcard.com/badge/github.com/rogerwelin/alviss" alt="Go Report Card"></a>
  <a href="https://github.com/rogerwelin/alviss/blob/master/LICENSE"><img src="https://img.shields.io/github/license/rogerwelin/alviss" alt="License"></a>
  <a href="https://github.com/rogerwelin/alviss/blob/master/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/rogerwelin/alviss" alt="Go version"></a>
</p>


**Alviss** is a scaffolding project that let's you provision and deploy production ready API:s in seconds on AWS using API Gateway and Lambda


Rationale
--------
Configuring API Gateway and Lambda using standard IaC tools like Terraform and Cloudformation is a very finicky and time consuming experience. Even with tools that are designed for serverless applications like *AWS SAM* and *serverless framework* can be hard and time consuming. Alviss is a scaffolder that takes care of the boilerplate using best practices and leaves you to tweak the settings. Use [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html) to finally deploy the project.


Installation
--------
Alviss is built in Go; meaning no runtime or dependencies to install, just grab a pre-built binary from the [GitHub Releases page](https://github.com/rogerwelin/alviss/releases). You can optionally put the **alviss** binary in your `PATH` so you can run alviss from any location.


Usage
--------

Example below shows how to generate a new public api using node.js as target language for the Lambda function(s):

```bash
$ ./alviss new-api -p my-project -e regional -l node
```

Then just follow the instructions on the screen. It's that simple!
