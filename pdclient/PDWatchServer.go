package pb

import "go.etcd.io/etcd/etcdserver/etcdserverpb"

type PDWatchServer struct {
	etcdserverpb.WatchServer
}

func NewPDWatchServer(ws etcdserverpb.WatchServer) PDWatchServer{
	return PDWatchServer{ws}
}