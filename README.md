# Meow Micro 🐈

Contains MicroServices which send and receive cat names. It is used
in the Distributed Tracing tutorial found here:

## Running Locally

In order to run locally simply perform the following:

1. Install Dependencies
    ```
    go install ...
    ```
2. Start the server
    ```
    $ go run server/server.go
    ```
3. Start the client
    ```
    $ go run client/client.go
    ```
4. Send a Request to the client-server:
    ```
    $ curl 127.0.0.1:5002 -X POST -d '{"body": "Meow-Mixer"}'
    ```

## Deploying to Kubernetes

I have included deployment scripts for deploying to Docker Desktop Kubernetes.

1. 
```
$ make clean
```

2. 
```
$ make build
```

3. 
```
$ make install
```

## Troubleshooting

1. Verify you can access the GRPC Service
```
grpcurl -d '{"body": "Meow-Mixer"}' -insecure localhost:443 chat.ChatService/SayHello

{
  "body": "Meow-Mixer"
}
```

## Tracing

Open Jaeger
