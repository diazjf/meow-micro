# Meow Micro 🐈

Contains MicroServices which send and receive cat names. The meow-client accepts requests via a REST API and uses GRPC to communicate with the meow-server.

It is used in the Distributed Tracing tutorial <name>. Which goes over:

- Ingress-Nginx Distributed Tracing
- Instrumenting MicroServices
- Viewing Trace

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

I have included deployment scripts for deploying to Docker-Desktop Kubernetes.
There are some pre-requisites required before running the below:

- Docker-Desktop
- Helm v3

1. Install Ingress
```
$ kubectl 
```

2. Install Jaeger-All-In-One
```
$ kubectl apply -f jaeger/
```

3. Update Ingress Config-Map
```
```

4. Build Docker Images
```
$ make build
```

5. Install via Helm Chart
```
$ make install
```

## Testing Deployment Tracing

1. Send a Request to the application
```
$ curl http://localhost/meow -X POST -d '{"name": "Meow-Mixer"}'
```

2. Verify GRPC communication via logs
```
$ kubectl logs

$ kubectl logs
```

3. Open Jaeger UI
```
$ chrome http://localhost:8081
```

4. See the Traces

## Troubleshooting

1. Check items are correctly deployed
```
$ kubectl get all

$ kubectl get all -n ingress-nginx
```

2. View the logs
```
$ kubectl logs

$ kubectl logs

$ kubectl logs
```

3. Exec into the meow-client pod
```
$ kubectl exec -it ....
```

4. install grpcurl
```
go get github.com/fullstorydev/grpcurl/...
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

5. Verify you can access the GRPC Service
```
grpcurl -d '{"body": "Meow-Mixer"}' -plaintext meow-server-svc:5001 chat.ChatService/SayHello

{
  "body": "Meow-Mixer"
}
```