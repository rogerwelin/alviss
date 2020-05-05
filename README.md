<p align="center">
  <a href="https://goreportcard.com/badge/github.com/rogerwelin/alviss"><img src="https://goreportcard.com/badge/github.com/rogerwelin/alviss" alt="Go Report Card"></a>
  <a href="https://github.com/rogerwelin/alviss/blob/master/LICENSE"><img src="https://img.shields.io/github/license/rogerwelin/alviss" alt="License"></a>
  <a href="https://github.com/rogerwelin/alviss/blob/master/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/rogerwelin/alviss" alt="Go version"></a>
</p>


**Alviss** is a scaffolding project that let's you provision and deploy production ready API:s in seconds on AWS using API Gateway and Lambda


Usage
--------

Example below shows how to generate a new public api using node.js as target language for the Lambda function(s):

```bash
$ ./alviss new-api -p my-project -e regional -l node
```

Then just follow the instructions on the screen. It's that simple!
