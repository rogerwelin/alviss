package main

// node.js
const nodeFunction = `
'use strict'

exports.handler = function(event, context) {

  let msg = {
    statusCode: 200,
    body: 'hello world'
  }
  context.succeed(msg);
};
`

const packageJSON = `
{
  "name": "helloworld",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "dependencies": {
  },
  "devDependencies": {
    "lambda-local": "^1.6.3"
  }
}
`

// python
const pythonFunction = `
import json

def handler(event, context):

    dict = { "msg" : "hello world" }
    
    return {
        'statusCode': 200,
        'body': json.dumps(dict)
    }
`

const requirementsFile = `
`

// ruby
const rubyFunction = `
require 'json'

def handler(event:, context:)
  msg = {:msg => 'hello world'}
  return {
    :statusCode => 200,
    :body => msg.to_json
  }
end
`
const gemFile = `
`

// go
const goFunction = `
package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type helloWorld struct {
	Msg string` + "`" + `json:"msg"` + "`" + `
}

func handler(ctx context.Context) (Response, error) {

	msg := helloWorld{Msg: "hello world"}
	resp, err := json.Marshal(msg)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	gwResp := Response{
		StatusCode: 200,
		Body:       string(resp),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return gwResp, nil
}

func main() {
	lambda.Start(handler)
}  
`

const goMod = `
module helloworld

go 1.13

require github.com/aws/aws-lambda-go v1.13.3
`
