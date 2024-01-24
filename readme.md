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
- first run
    - cd xray && docker build -t xray-daemon .
- demo
    1. 確定 account_id && region, `make show`
    2. `docker-compose up`
    3. `cd ./app/utility/ && go test -run gorm_test.go`
    4. `make test`

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

## apis with database, by Gorm
> 目前 xray 使用 `xray.SQLContext`, 不支援 gorm, 只能用繞路的方式.
> 且, xray-console 的 map 無法顯示 database 的 node.

- `curl -X POST $server:9001/new`
- `curl -X GET $server:9001/all`
- `curl -X DELETE $server:9001/del -d '{"id":1}'`

## apis with databse, by sql.DB
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