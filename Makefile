.PHONY: xray app

AWS_ACCOUNT_ID="424613967558"
AWS_REGION="ap-southeast-1"

app:
	@aws ecr get-login-password \
		--region $(AWS_REGION) \
		| docker login \
		--username AWS \
		--password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com;
	@docker run --rm \
		-v $$(pwd)/app:/home/app \
		-w /home/app golang:1.14.3 \
		go build;
	@cd app;\
	docker build -t $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/xray-app;\
	docker push $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/xray-app

xray:
	@aws ecr get-login-password \
		--region $(AWS_REGION) \
		| docker login \
		--username AWS \
		--password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com;
	@cd xray;\
	docker build -t $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/aws-xray-daemon .;\
	docker push $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/aws-xray-daemon