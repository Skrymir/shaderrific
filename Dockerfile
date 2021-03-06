FROM golang:1.14-alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

COPY config.yml /go/bin

WORKDIR /go/bin
CMD ["shaderrific"]
