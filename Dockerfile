FROM golang:1.15-alpine as stage1
ENV SRC_DIR=/go/src/github.com/ajarv/go-app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /work
ENV GO111MODULE=on

COPY . .
RUN go mod download &&\
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-web ./cmd/web/main.go

FROM alpine:latest  

RUN apk update && apk upgrade && \
    apk --no-cache add ca-certificates curl bash && \
    mkdir -p /work 

COPY --from=stage1 /work/go-web /work/go-web

USER 1000
EXPOSE 8080
EXPOSE 8443
CMD /work/go-web