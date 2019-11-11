package main

const swagger = `
---
openapi: "3.0.1"
{{ if and (eq .apiEndpoints "private") }}
x-amazon-apigateway-policy:
  Version: '2012-10-17'
  Statement:
    - Effect: Allow
      Principal: "*"
      Action: 
        - "execute-api:Invoke"
      Resource: "execute-api:/*"
      Condition:
        StringEquals:
          aws:SourceVpc:
            Ref: VPCId
{{ else }}
{{ end }}

info:
  title: {{ .apiProjectName }}
  description: your awesome description here
  version: "v1.0"

servers:
  - url: https://apigw-url.example.com/prod
    description: Test environment URL
  - url: http://apigw-url.example.com/test
    description: Production environment URL

paths:
  /v1/{{ .lambdaFunctionName }}/{userId}:
    get:
      summary: hello world endpoint
      description: outputs hello world
      parameters:
        - in: path
          name: userId
          schema:
            type: integer
          required: true
      responses:
        200:
          description: "OK"
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/ArticleObj"
        500:
          description: "Internal Server Error"
          content: {}
      x-amazon-apigateway-integration:
        uri:
          Fn::Sub: "arn:${AWS::Partition}:apigateway:${AWS::Region}:lambda:path/2015-03-31/functions/${HelloWorldFunction.Arn}/invocations"
        httpMethod: POST
        passthroughBehavior: "when_no_match"
        type: aws_proxy
components:
  schemas:
    ArticleObj:
      properties:
        rowValues:
          type: "array"
          items:
            type: "object"
            properties:
              modelId:
                type: "number"
              Article_Id:
                type: "string"
              MU_UOM_Cd:
                type: "string"
              rank:
                type: "number"
`
