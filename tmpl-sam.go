package main

const apiGWConf = `
---
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31

{{ if and (eq .APIEndpoints "private") }}
Parameters:
  VPCId:
    Description: ID of the VPC ID
    Type: AWS::EC2::VPC::Id
  VPCCidr:
    Description: Valid CIDR of the VPC
    Type: String
    AllowedPattern: ^([0-9]{1,3}\.){3}[0-9]{1,3}(\/([0-9]|[1-2][0-9]|3[0-2]))?$
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x
  SubnetIDs:
    Description: A list/array of Subnet IDs
    Type: List<AWS::EC2::Subnet::Id>
  Environment:
    Description: name of the environment
    Type: String
    Default: test
    AllowedValues: [test, prod]
{{ else }}
Parameters:
  Environment:
    Description: name of the environment
    Type: String
    Default: test
    AllowedValues: [test, prod]
{{ end}}

Conditions:
  IsProd:
    !Equals [!Ref Environment, "prod"]

Resources:

{{ if and (eq .APIEndpoints "private") }}
  ########################
  # Infra stuff
  ########################

  LambdaSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: SG for private lambfa functions
      GroupName: vpc-lambda
      VpcId: !Ref VPCId
      SecurityGroupIngress:
        - IpProtocol: tcp
          CidrIp: !Ref VPCCidr
          FromPort: 0
          ToPort: 65535

  ApiGwVpcEndpoint:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      ServiceName: !Sub "com.amazonaws.${AWS::Region}.execute-api"
      PrivateDnsEnabled: true
      VpcEndpointType: Interface
      VpcId: !Ref VPCId
      SubnetIds: !Ref SubnetIDs
{{ end }}

  ########################
  # API GW Conf
  ########################

  AWSApi:
    Type: AWS::Serverless::Api
    Properties:
      StageName: !Ref Environment
      TracingEnabled: true # Enable X-Ray for distributed tracing to help debugging
      {{ if and (eq .APIEndpoints "private")}}EndpointConfiguration: PRIVATE{{ end }}{{ if and (eq .APIEndpoints "regional")}}EndpointConfiguration: REGIONAL{{ end }}{{ if and (eq .APIEndpoints "edge")}}EndpointConfiguration: EDGE{{ end }}
      # Use DefinitionBody for swagger file so that we can use CloudFormation functions within the swagger file
      DefinitionBody:
        'Fn::Transform':
          Name: 'AWS::Include'
          Parameters:
            Location: ./swagger-api.yml
      MethodSettings:
        - ResourcePath: '/*'
          HttpMethod: '*'
          LoggingLevel: INFO
          MetricsEnabled: true    # Enable detailed metrics
          DataTraceEnabled: true  # Put logs into cloudwatch
      AccessLogSetting:
        DestinationArn: !Sub "arn:${AWS::Partition}:logs:${AWS::Region}:${AWS::AccountId}:log-group:${ApiAccessLogGroup}"
        Format: '$context.identity.sourceIp $context.authorizer.claims.sub [$context.requestTime] "$context.httpMethod $context.resourcePath $context.protocol" $context.status $context.requestId $context.awsEndpointRequestId $context.xrayTraceId $context.responseLatency $context.integrationLatency "$context.error.message"'
      Cors:
        AllowOrigin: "'*'"
        AllowHeaders: "'content-type'"

  ########################
  # IAM for API GW
  ########################

  # This role allows API Gateway to push execution and access logs to CloudWatch logs
  ApiGatewayPushToCloudWatchRole:
    Type: "AWS::IAM::Role"
    Properties:
      Description: "Push logs to CloudWatch logs from API Gateway"
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: Allow
            Principal:
              Service: apigateway.amazonaws.com
            Action: "sts:AssumeRole"
      ManagedPolicyArns:
        - !Sub "arn:${AWS::Partition}:iam::aws:policy/service-role/AmazonAPIGatewayPushToCloudWatchLogs"

  ApiAccount:
    Type: "AWS::ApiGateway::Account"
    Properties:
      CloudWatchRoleArn: !GetAtt ApiGatewayPushToCloudWatchRole.Arn

  ApiAccessLogGroup:
    Type: AWS::Logs::LogGroup
    Properties:
      LogGroupName: !Sub "/aws/apigateway/AccessLog-${AWSApi}"
      RetentionInDays: 365


  ########################
  #  Functions Goes Here
  ########################

  HelloWorldFunction:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: src/{{ .LambdaFunctionName }}/app/
      {{ if and (eq .Language "python")}}Handler: app.handler
      Runtime: python3.7{{ end }}{{ if and (eq .Language "ruby")}}Handler: app.handler
      Runtime: ruby2.7{{ end }}{{ if and (eq .Language "node")}}Handler: index.handler
      Runtime: nodejs12.x{{ end }}{{ if and (eq .Language "go")}}Handler: helloworld
      Runtime: go1.x{{ end }}{{ if and (eq .Language "java")}}Handler: com.api.HelloWorld
      Runtime: java8{{ end }}
      MemorySize: 512
      Timeout: 5
      Tracing: Active
      Policies:
{{ if and (eq .APIEndpoints "private") }}        - AWSLambdaVPCAccessExecutionRole{{ else }}        - AWSLambdaExecute{{ end }}
      Events:
        AnyApi:
          Type: Api
          Properties:
            RestApiId: !Ref AWSApi
            Path: '/{{ .LambdaFunctionName }}'
            Method: GET
{{ if and (eq .APIEndpoints "private") }}      VpcConfig:
        SecurityGroupIds:
          - !GetAtt LambdaSecurityGroup.GroupId
        SubnetIds: !Ref SubnetIDs
{{ end }}

Outputs:
  ApiURL:
    Description: {{ .APIProjectName }}
    Value: !Sub 'https://${AWSApi}.execute-api.${AWS::Region}.amazonaws.com/${Environment}/'
`
