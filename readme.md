# readme
```bash
.
├── app/                        <-- golang app, restful api with `aws-xray-sdk-go`
├── xray/                       <-- xray Dockerfile, localhost log
├── docker-compose.yml          <-- localhost build Sample Applications
├── Makefile                    <-- docker push app/xray to ECR
└── cfn.yml                     <-- AWS CloudFormation Template
```

# pre-required
- Local AWS Credential

# run localhost
- run 
    1. `docker-compose up`
    2. `curl localhost:9001`

![img](./assets/localhost.png)


- go to aws xray service to see the information (wait 1 min)

![img](./assets/xray.png)

# deploy to ecs-fargate, steps

1. go to ECR and create three repositories
    - xray-app
    - xray-mariadb
    - aws-xray-daemon
2. edit Makefile, to update the below args
    - AWS_ACCOUNT_ID
    - AWS_REGION
3. push image to ECR
    - `make app`
    - `make xray`
    - `make mysql`
4. go to AWS CloudFormation and build stack with `cfn.yml`
5. go to ECS console and get Task public IP, `curl $IP` and then check XRAY console


# xray data generated
- Localhost test, use `make test` to execute below command
- `server=localhost`
## base
- `curl $server:9001`
- `curl $server:9001/ping`
## apis with database 
- `curl -X POST $server:9001/new`
- `curl -X GET $server:9001/all`
- `curl -X DELETE $server:9001/del -d '{"id":1}'`
- `curl -X GET $server:9001/sql/by/xray/success`
- `curl -X GET $server:9001/sql/by/xray/error`
## err
- `curl -X GET $server:9001/error/400`
- `curl -X GET $server:9001/error/429`
- `curl -X GET $server:9001/error/500`
- `curl -X GET $server:9001/error/panic`
## many funcs
- `curl -X PATCH $server:9001/many/funcs`
- `curl -X PATCH $server:9001/send/sqs`