FROM golang:1.16

WORKDIR /go/src/app
COPY ../. /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 5001
EXPOSE 5002

CMD ["client"]
