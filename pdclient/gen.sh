GOGO_ROOT=${GOPATH}/src/github.com/gogo/protobuf

protoc -I.:${GOPATH}/src --proto_path=./../etcdserver/etcdserverpb/rpc.proto --gogo_out=plugins=grpc:. pd.proto
