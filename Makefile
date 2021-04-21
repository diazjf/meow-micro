.PHONY: clean
clean-image:
	docker rmi -f meow-client:1.0
	docker rmi -f meow-server:1.0

.PHONY: build
build:
	docker build -t meow-client:1.0 -f client/Dockerfile .
	docker build -t meow-server:1.0 -f server/Dockerfile .

.PHONY: push
push:
	docker push meow-client:1.0
	docker push meow-server:1.0
