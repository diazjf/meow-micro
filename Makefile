CLIENT_IMAGE = meow-client:1.0
SERVER_IMAGE = meow-server:1.0
HELM_DEPLOY = meow-micro
SECRET_NAME = grpc-secret

.PHONY: clean
clean:
	docker rmi -f $(CLIENT_IMAGE)
	docker rmi -f $(SERVER_IMAGE)
	helm delete $(HELM_DEPLOY)
	rm -rf tmp
	kubectl delete secret $(SECRET_NAME)

.PHONY: build
build:
	docker build -t $(CLIENT_IMAGE) -f client/Dockerfile .
	docker build -t $(SERVER_IMAGE) -f server/Dockerfile .

.PHONY: install
install:
	mkdir tmp
	openssl req -x509 -nodes -newkey rsa:2048 -days 365 -keyout /tmp/localhost.key -out /tmp/localhost.crt -subj "/CN=localhost/O=localhost"
	kubectl create secret tls $(SECRET_NAME) --key /tmp/localhost.key --cert /tmp/localhost.crt
	helm install meow-micro helm