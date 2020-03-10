### 이화일은 바로 실행하지 말고 dkcu 연결해서 빌드할 것 ###

FROM golang:1.11-alpine 

# RUN  apt-get update
# RUN  apt-get install git-core

# RUN yum install git
# alpine version 은 apk로 설치함.
RUN apk add git

RUN mkdir -p /go/src/github.com/dabory/svc-abango
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN go get xorm.io/xorm 
RUN go get github.com/pilu/fresh
RUN go get github.com/go-sql-driver/mysql

RUN  go get github.com/dabory/abango 

WORKDIR  /go/src/github.com/dabory/svc-abango
ADD . /go/src/github.com/dabory/svc-abango