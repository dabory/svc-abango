# gRPC KangAn How to

## Running the server and client

Note : You need golang, node and yarn installed for this. And the server runs at port `8081`. It's a hard coded port number in the code

To run the server, use this command

#### Step 0: Linux: Ubuntu 16.04.4 LTS \n \l  ; recommended Linux version
#### Step 1: Install: Go version go1.11 linux/amd64
#### Step 2: Install: Go libraries
`$ go get -u google.golang.org/grpc`

`$ go get -u github.com/golang/protobuf/protoc-gen-go`

`$ go get xorm.io/xorm`

`$ go get github.com/pilu/fresh`

`$ go get github.com/go-sql-driver/mysql`

`$ go get github.com/dabory/abango `

`$ mkdir -p $GOPATH/bin $GOPATH/src $GOPATH/pkg`

`$ cd $GOPATH/src`

`$ git clone http://13.124.244.187:10080/kangan/github.com/dabory/svc-abango `

To run the grpc server, using this command

`$ cd $GOPATH/src/github.com/dabory/svc-abango`

#### Step 3: Change Open port if you need in conf/ folder files
`$ vi conf/run_conf.json`

`$ vi conf/xxx_conf.json`  ; liked conf file

#### Step 5: Create MySQL Tables : extract table schema from kangan_db-191125.sql 

#### Step 4: Run go lang server
`$ fresh`
