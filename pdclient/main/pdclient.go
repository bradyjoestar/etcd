package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pdclient"
	"google.golang.org/grpc"
	"log"
	"math"
)


func main() {
	conn, err := grpc.Dial("localhost:28080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("faild to connect: %v", err)
	} else {
		fmt.Println("connection success")
	}
	defer conn.Close()

	watchClient := etcdserverpb.NewWatchClient(conn)
	fmt.Println("watchclient success")

	watcher := clientv3.NewWatcherWithCallOption(watchClient,[]grpc.CallOption{
		grpc.FailFast(false),
		grpc.MaxCallSendMsgSize(2 * 1024 * 1024),
		grpc.MaxCallRecvMsgSize(math.MaxInt32),
	})

	pdwatcher := pb.NewPDWatcher(watcher)

	for{
		rch := pdwatcher.Watch(context.Background(), "mykey")
		for wresp := range rch {
			fmt.Println(wresp.Header)
			fmt.Println("we have receive new event")
		}
	}

}