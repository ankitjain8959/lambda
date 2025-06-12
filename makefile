APP_NAME = LambdaFunctionUsingGo
LAMBDA_PAYLOAD = resources/event.json

.PHONY: build clean local-invoke

build:
	@echo "Building Lambda..."
	sam build

clean:
	PowerShell -Command "Remove-Item -Recurse -Force .aws-sam\\build"

run:
# sam local invoke $(APP_NAME) --event $(LAMBDA_PAYLOAD)
	sam local invoke --docker-network lambda-mongo-net $(APP_NAME) --event $(LAMBDA_PAYLOAD)