package connect

import "go.etcd.io/etcd/clientv3"

var (
	etcdClient *clientv3.Client
	etcdConf   *etcdConfig
)

type etcdConfig struct {
	endpoints []string
	user      string
	password  string
	etcd      *clientv3.Client
}

func Etcd() *clientv3.Client {
	if etcdClient == nil {
		if etcdConf.etcd == nil {
			_ = etcdConf.Connect()
		}
		etcdClient = etcdConf.etcd
	}
	return etcdClient
}

func NewEtcdConfig(endpoints []string, user, password string) *etcdConfig {
	etcdConf = &etcdConfig{
		endpoints: endpoints,
		user:      user,
		password:  password,
		etcd:      nil,
	}
	return etcdConf
}

func (e *etcdConfig) Connect() (err error) {
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints: e.endpoints,
		Username:  e.user,
		Password:  e.password,
	})
	e.etcd = etcdClient
	return nil
}

func (e *etcdConfig) Close() error {
	return e.etcd.Close()
}
