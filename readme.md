# readme
```bash
.
├── app/                        <-- golang app, restful api with `aws-xray-sdk-go`
├── xray/                       <-- xray Dockerfile, localhost log
├── docker-compose.yml          <-- localhost use xray
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

1. go to ECR and create two repositories
    - xray-app
    - aws-xray-daemon
2. go to Makefile, to update the args
    - AWS_ACCOUNT_ID
    - AWS_REGION
3. push image to ECR
    - `make app`
    - `make xray`
4. go to AWS CloudFormation and build stack with `cfn.yml`
5. go to ECS console and get Task public IP, `curl $IP` and then check XRAY console



