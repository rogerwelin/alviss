package main

const apiGWConf = `
---
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31

{{ if and (eq .ApiEndpoints "private") }}
Parameters:
  VPCId:
    Description: ID of the VPC ID
    Type: AWS::EC2::VPC::Id
  SubnetIDs:
    Description: A list/array of Subnet IDs
    Type: List<AWS::EC2::Subnet::Id>
  Environment:
    Description: name of the environment
    Type: String
    AllowedValues: [test, prod]
{{ else }}
Parameters:
  Environment:
    Description: name of the environment
    Type: String
    AllowedValues: [test, prod]
{{ end}}

Conditions:
  IsProd:
    !Equals [!Ref Environment, "prod"]

Resources:

{{ if and (eq .ApiEndpoints "private") }}
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
          CidrIp: 172.31.190.0/24
          FromPort: 0
          ToPort: 65535
        - IpProtocol: tcp
          CidrIp: 172.31.178.0/23
          FromPort: 0
          ToPort: 65535
        - IpProtocol: tcp
          CidrIp: 10.168.58.0/24
          FromPort: 0
          ToPort: 65535
        - IpProtocol: tcp
          CidrIp: 172.31.176.0/23
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

  AnalyticsApi:
    Type: AWS::Serverless::Api
    Properties:
      StageName: !Ref Environment
      TracingEnabled: true # Enable X-Ray for distributed tracing to help debugging
      {{ if and (eq .ApiEndpoints "private")}}EndpointConfiguration: PRIVATE{{ end }}{{ if and (eq .ApiEndpoints "regional")}}EndpointConfiguration: REGIONAL{{ end }}{{ if and (eq .ApiEndpoints "edge")}}EndpointConfiguration: EDGE{{ end }}
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
      LogGroupName: !Sub "/aws/apigateway/AccessLog-${AnalyticsApi}"
      RetentionInDays: 365


  ########################
  #  Functions Goes Here
  ########################

  RecommendationsFunction:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: src/recommendations/app/
      Handler: recommendations.lambda_handler
      Runtime: python3.7
      MemorySize: 512
      Timeout: 5
      Tracing: Active
      Policies:
        - AWSLambdaExecute
      Layers:
        - !Ref RecommendationsLayer
      Events:
        AnyApi:
          Type: Api
          Properties:
            RestApiId: !Ref AnalyticsApi
            Path: '/recommendations/{userId}'
            Method: GET

  RecommendationsLayer:
    Type: AWS::Serverless::LayerVersion
    Properties:
      LayerName: recommendations-deps
      Description: Dependencies for RecommendationsFunction
      ContentUri: src/recommendations/dependencies/
      CompatibleRuntimes:
        - python3.7
      RetentionPolicy: Retain

Outputs:
  ApiURL:
    Description: {{ .ApiProjectName }}
    Value: !Sub 'https://${AnalyticsApi}.execute-api.${AWS::Region}.amazonaws.com/${Environment}/'
`
