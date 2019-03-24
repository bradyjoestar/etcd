package pb

import "go.etcd.io/etcd/clientv3"

type PDWatcher struct {
	clientv3.Watcher
}

func NewPDWatcher(watcher clientv3.Watcher) PDWatcher{
	return PDWatcher{watcher}
}