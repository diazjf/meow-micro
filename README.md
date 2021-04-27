# Meow Micro 🐈

Contains MicroServices which send and receive cat names. The meow-client accepts requests via a REST API and uses GRPC to communicate with the meow-server.

It is used in the Distributed Tracing tutorial <name>. Which goes over:

- Ingress-Nginx Distributed Tracing
- Instrumenting MicroServices
- Viewing Traces

## Running Locally

In order to run locally simply perform the following:

1. Install Dependencies
    ```
    $ go install ...
    ```
2. Start the server
    ```
    $ go run server/server.go
    ```
3. Start the client in another terminal
    ```
    $ go run client/client.go
    ```
4. Send a Request to the client
    ```
    $ curl 127.0.0.1:5002 -X POST -d '{"body": "Meow-Mixer"}'
    ```
5. View the output of both in the terminal

## Deploying to Kubernetes

I have included deployment scripts for deploying to Docker-Desktop Kubernetes.
There are some pre-requisites required before running the below:

- Docker-Desktop
- Helm v3

1. Install Ingress
```
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.45.0/deploy/static/provider/cloud/deploy.yaml
```

2. Install Jaeger-All-In-One
```
$ kubectl apply -f jaeger/jaeger-all-in-one.yaml
```

3. Update Ingress Config-Map
```
$ echo '
  apiVersion: v1
  kind: ConfigMap
  data:
    enable-opentracing: "true"
    jaeger-collector-host: jaeger-agent.default.svc.cluster.local              
  metadata:
    name: ingress-nginx-controller
    namespace: ingress-nginx
  ' | kubectl replace -f -
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

2. Open Jaeger UI
```
$ open http://localhost:8081
```

3. See the Traces

## Troubleshooting

1. Check items are correctly deployed
```
$ kubectl get all
$ kubectl get all -n ingress-nginx
```

2. View the logs
```
$ kubectl logs -n ingress-nginx <ingress-controller>
$ kubectl logs <meow-server>
$ kubectl logs <meow-client>
```

3. Exec into the meow-client pod
```
$ kubectl exec -it <meow-client> -- bash
```

4. install grpcurl
```
$ go get github.com/fullstorydev/grpcurl/...
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

5. Verify you can access the GRPC Service
```
$ grpcurl -d '{"body": "Meow-Mixer"}' -plaintext meow-server-svc:5001 chat.ChatService/SayHello
```