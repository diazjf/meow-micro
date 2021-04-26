CLIENT_IMAGE = meow-client:1.0
SERVER_IMAGE = meow-server:1.0
HELM_DEPLOY = meow-micro
SECRET_NAME = grpc-secret

.PHONY: clean
clean:
	docker rmi -f $(CLIENT_IMAGE)
	docker rmi -f $(SERVER_IMAGE)
	helm delete $(HELM_DEPLOY)

.PHONY: build
build:
	docker build -t $(CLIENT_IMAGE) -f client/Dockerfile .
	docker build -t $(SERVER_IMAGE) -f server/Dockerfile .

.PHONY: install
install:
	helm install meow-micro helm