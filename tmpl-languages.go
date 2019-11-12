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
