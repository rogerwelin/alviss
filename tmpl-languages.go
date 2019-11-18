package main

// node.js
const nodeFunction = `
'use strict'
const  winston = require('winston')

const logger = winston.createLogger({
  level: 'info',
  format: winston.format.combine(
    winston.format.timestamp({
      format: 'YYYY-MM-DD HH:mm:ss'
    }),
  ),
});

exports.handler = function(event, context) {
  logger.log({
    level: 'info',
    message: event,
  });

  let msg = {
    msg: 'hello world'
  }
  context.succeed(JSON.stringify(msg));
};
`

const packageJson = `
{
  "name": "hello",
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
    "winston": "^3.2.1"
  },
  "devDependencies": {
    "lambda-local": "^1.6.3"
  }
}
`

// python
const pythonFunction = `
import json

def lambda_handler(event, context):

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
    { statusCode: 200, body: JSON.generate('hello world') }
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

func Handler(ctx context.Context) (Response, error) {

	msg := helloWorld{Msg: "hwllo world"}
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
	lambda.Start(Handler)
}  
`

const goMod = `
module helloworld

go 1.13

require github.com/aws/aws-lambda-go v1.13.3
`
