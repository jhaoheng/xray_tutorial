.PHONY: xray app mysql test

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
	docker build -t $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/xray-app .;\
	docker push $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/xray-app

mysql:
	@aws ecr get-login-password \
		--region $(AWS_REGION) \
		| docker login \
		--username AWS \
		--password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com;
	@docker pull mariadb:10.5.3;\
	docker tag mariadb:10.5.3 $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/xray-mariadb:10.5.3;\
	docker push $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/xray-mariadb:10.5.3;

xray:
	@aws ecr get-login-password \
		--region $(AWS_REGION) \
		| docker login \
		--username AWS \
		--password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com;
	@cd xray;\
	docker build -t $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/aws-xray-daemon .;\
	docker push $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/aws-xray-daemon

test: 
	@set -x;\
	which curl;\
	if [ "0" -ne "$$?" ];then\
		echo "For test, 'curl' command needed!!!!";\
		exit 0;\
	fi;\
	curl -s localhost:9001 >> /dev/null;\
	curl -s localhost:9001/ping >> /dev/null;\
	curl -s -X POST localhost:9001/new >> /dev/null;\
	curl -s -X GET localhost:9001/all >> /dev/null;\
	curl -s -X DELETE localhost:9001/del -d '{"id":1}' >> /dev/null;\
	curl -s -X GET localhost:9001/error/400 >> /dev/null;\
	curl -s -X GET localhost:9001/error/429 >> /dev/null;\
	curl -s -X GET localhost:9001/error/500 >> /dev/null;\
	curl -s -X GET localhost:9001/error/panic >> /dev/null;\
	curl -s -X PATCH localhost:9001/many/funcs >> /dev/null;\
	curl -s -X PATCH localhost:9001/send/sqs >> /dev/null;\
	set +x;